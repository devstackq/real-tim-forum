package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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

func (us *UserService) Create(user *models.User) (int, int, error) {
	//good practice ?
	validEmail := IsEmailValid(user)
	validPassword := IsPasswordValid(user)

	if validEmail && validPassword {
		// if utils.AuthType == "default" {
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			return http.StatusInternalServerError, -1, err
		}
		user.Password = string(hashPwd)
		//go to repo, interface -> method call
		lastId, err := us.repository.CreateUser(user)
		//check  is already user
		if err != nil {
			if sqliteErr, ok := err.(sqlite.Error); ok {
				if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
					return http.StatusBadRequest, -1, errors.New("nickname or email exist")
				}
			}
			return http.StatusInternalServerError, -1, err
		}
		return http.StatusOK, int(lastId), nil
	} else if !validEmail {
		return http.StatusBadRequest, 0, errors.New("email incorrect, example@mail.com")
	} else {
		return http.StatusBadRequest, 0, errors.New("password incorrect, 123User!")
	}
}

func (us *UserService) Logout(session string) error {
	err := us.repository.Logout(session)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserById(id string) (*models.User, error) {
	user, err := us.repository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (us *UserService) GetCreatedUserPosts(userId int) (*[]models.Post, error) {
	posts, err := us.repository.GetCreatedUserPosts(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (us *UserService) GetUserVotedItems(userId int) (*[]models.Vote, error) {
	voted, err := us.repository.GetUserVotedItems(userId)
	if err != nil {
		return nil, err
	}
	return voted, nil
}

// take out utils ? user_utils.
