package service

import (
	"errors"
	"net/http"
	"regexp"
	"time"
	"unicode"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	sqlite "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.User
}

//repoUser wrapper struct - userService
func NewUserService(repo repository.User) *UserService {
	return &UserService{repo}
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

		user.Password = string(hashPwd)
		user.CreatedTime = time.Now()

		lastId, err := us.repo.CreateUser(user)

		if err != nil {
			if sqliteErr, ok := err.(sqlite.Error); ok {
				if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
					return http.StatusBadRequest, -1, errors.New("User already created")
				}
			}
			return http.StatusInternalServerError, -1, err
		}
		return http.StatusOK, int(lastId), nil
	} else {
		//else - > send message json, error -> exist email || pwd
		return http.StatusBadRequest, 0, errors.New("email or pwd exist")
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
