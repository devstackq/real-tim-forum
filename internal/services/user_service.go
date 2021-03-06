package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.User
}

//repoUser wrapper struct - userService
func NewUserService(repo repository.User) *UserService {
	fmt.Println("create user service ")
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
	log.Println("session update -> signin in system", session.UUID)
	return http.StatusOK, &session, nil
}

func (us *UserService) CreateUser(user *models.User) (int, int, error) {
	//good practice ?
	validEmail := IsEmailValid(user)
	validPassword := IsPasswordValid(user)

	// validFields = [""]
	if IsEmpty(user.FullName) {
		return http.StatusBadRequest, 0, errors.New("empty full name field")
	} else if IsEmpty(user.Age) {
		return http.StatusBadRequest, 0, errors.New("empty age field")
	} else if IsEmpty(user.City) {
		return http.StatusBadRequest, 0, errors.New("empty city field")
	} else if IsEmpty(user.Sex) {
		return http.StatusBadRequest, 0, errors.New("empty sex field")
	} else if IsEmpty(user.Username) {
		return http.StatusBadRequest, 0, errors.New("empty nick name field")
	} else if !validEmail {
		return http.StatusBadRequest, 0, errors.New("email incorrect, example@mail.com")
	} else if !validPassword {
		return http.StatusBadRequest, 0, errors.New("password incorrect")
	} else {
		if validEmail && validPassword {
			hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
			if err != nil {
				return http.StatusInternalServerError, -1, err
			}
			user.Password = string(hashPwd)
			//go to repo, interface -> method call
			_, err = us.repository.CreateUser(user)
			//check  is already user
			if err != nil {
				return http.StatusBadRequest, -1, errors.New("nickname or email exist")
						}
		}
	}
	return http.StatusOK, 0, nil

}

func (us *UserService) Logout(session string) error {
	err := us.repository.Logout(session)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserUUID(userId int) (string, error) {
	uuid, err := us.repository.GetUserUUID(userId)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (us *UserService) GetUserID(uuid string) (int, error) {
	userid, err := us.repository.GetUserID(uuid)
	if err != nil {
		return 0, err
	}
	return userid, nil
}

func (us *UserService) GetUserName(uuid string) (string, error) {
	name, err := us.repository.GetUserName(uuid)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (us *UserService) GetUserProfile(userId int) (*models.Profile, error) {
	data, err := us.repository.GetProfileData(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
