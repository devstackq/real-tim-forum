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

func (us *UserService) Signin(user *models.User) (int, *models.Session, error) {

	id, hashPassword, err := us.repository.SigninUser(user)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password))
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	fmt.Println("pwd and login corect", id, user)

	uuid := uuid.Must(uuid.NewV4(), err).String()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	session := models.Session{}
	session.UserID = id
	session.UUID = uuid
	session.StartTimeCookie = time.Now()

	err = us.repository.UpdateSession(&session)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, &session, nil

	//SetCookie in Browser

	//1 time set uuid user, set cookie in Browser

	//event create post/comment, etc -> middleware()
	//middleware -> get userId & email(from client) -> chek in DB -> have session table
	//uuid field -> compare uuidDb & uuid localStorage -> if true -> run handler

}

func (us *UserService) Create(user *models.User) (int, int, error) {

	validEmail := us.isEmailValid(user)
	validPassword := us.isPasswordValid(user)

	if validEmail && validPassword {
		// if utils.AuthType == "default" {
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			return http.StatusInternalServerError, -1, err
		}
		user.Password = string(hashPwd)
		user.CreatedTime = time.Now()
		//go to repo, interface -> method call
		lastId, err := us.repository.CreateUser(user)
		//check  is already user
		if err != nil {
			if sqliteErr, ok := err.(sqlite.Error); ok {
				if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
					return http.StatusBadRequest, -1, errors.New("Nickname or email exist")
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

func (us *UserService) GetDataInDb(uid string, what string) (string, error) {

	var data string
	var err error

	if what == "uuid" {
		data, err = us.repository.GetUuidInDb(uid)
		if err != nil {
			return "", err
		}
	}
	return data, nil
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
