package repository

import (
	"database/sql"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

//have db - connect
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

//implement method, by interface User
func (ur *UserRepository) CreateUser(user *models.User) (int64, error) {

	query, err := ur.db.Prepare(`INSERT INTO users(full_name, email, username, password, age, sex, created_time, city, image) VALUES(?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return -1, err
	}
	result, err := query.Exec(user.FullName, user.Email, user.Username, user.Password, user.Age, user.Sex, time.Now(), user.City, user.Image)
	if err != nil {
		return -1, err
	}
	defer query.Close()

	return result.LastInsertId()
}
