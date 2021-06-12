package service

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"
	"unicode"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	sqlite "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.User
}

//repoUser wrapper struct - userService
func NewUserService(repo repository.User) *UserService {
	return &UserService{repo}
}

func (us *UserService) Signin(user *models.User) (int, error) {

	id, hashPassword, err := us.repository.SigninUser(user)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password))
	if err != nil {
		return http.StatusBadRequest, err
	}
	fmt.Println("pwd and login corect", id, user)

	uuid := uuid.Must(uuid.NewV4(), err).String()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	session := models.Session{}
	session.UserID = id
	session.UUID = uuid
	session.StartTimeCookie = time.Now()

	us.repository.UpdateSession(&session)

	//SetCookie in Borowser

	// //1 time set uuid user, set cookie in Browser
	// newSession := general.Session{
	// 	UserID: user.ID,
	// }

	// uuid := uuid.Must(uuid.NewV4(), err).String()
	// if err != nil {
	// 	utils.AuthError(w, r, err, "uuid problem", utils.AuthType)
	// 	return
	// }
	// //create uuid and set uid DB table session by userid,
	// userPrepare, err := DB.Prepare(`INSERT INTO session(uuid, user_id, cookie_time) VALUES (?, ?, ?)`)
	// if err != nil {
	// 	log.Println(err)
	// }

	// _, err = userPrepare.Exec(uuid,  newSession.UserID, time.Now())
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer userPrepare.Close()

	// if err != nil {
	// 	utils.AuthError(w, r, err, "the user is already in the system", utils.AuthType)
	// 	//get ssesion id, by local struct uuid
	// 	log.Println(err)
	// 	return
	// }

	// // get user in info by session Id
	// err = DB.QueryRow("SELECT id, uuid FROM session WHERE user_id = ?", newSession.UserID).Scan(&newSession.ID, &newSession.UUID)
	// if err != nil {
	// 	utils.AuthError(w, r, err, "not find user from session", utils.AuthType)
	// 	log.Println(err, "her")
	// 	return
	// }
	// utils.SetCookie(w, newSession.UUID)

	return http.StatusOK, nil

}

func (us *UserService) Create(user *models.User) (int, int, error) {
	//create user, use middleware, chek email, pwd another things
	// -> then call repos.CreateUser()
	validEmail := us.isEmailValid(user)
	validPassword := us.isPasswordValid(user)

	if validEmail && validPassword {
		// if utils.AuthType == "default" {
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			return http.StatusInternalServerError, -1, err
		}
		fmt.Println(user.Username, "userNmae")
		user.Password = string(hashPwd)
		user.CreatedTime = time.Now()
		//go to repo, interface -> method call
		lastId, err := us.repository.CreateUser(user)
		//check  is already user
		if err != nil {
			fmt.Println(err)
			if sqliteErr, ok := err.(sqlite.Error); ok {
				fmt.Println(sqliteErr)
				if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
					return http.StatusBadRequest, -1, errors.New("User already created")
				}
			}
			return http.StatusInternalServerError, -1, err
		}
		fmt.Println(user, "Create user service done")

		return http.StatusOK, int(lastId), nil

	} else if !validEmail {
		return http.StatusBadRequest, 0, errors.New("email incorrect, example@mail.com")
	} else {
		return http.StatusBadRequest, 0, errors.New("password incorrect, 123User!")
	}
}

func (us *UserService) isEmailValid(user *models.User) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,6}$`)
	return Re.MatchString(user.Email)
}

func (us *UserService) isPasswordValid(user *models.User) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(user.Password) >= 7 {
		hasMinLen = true
	}
	for _, char := range user.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
