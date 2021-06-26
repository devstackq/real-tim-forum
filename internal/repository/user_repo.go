package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
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

	query, err := ur.db.Prepare(`INSERT INTO users(full_name, email, user_name, password, age, sex, created_time, city, image) VALUES(?,?,?,?,?,?,?,?,?)`)
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

func (ur *UserRepository) SigninUser(user *models.User) (int, string, error) {

	var id int
	var hashPassword string

	sqlStatement := `SELECT id, password FROM users WHERE email=?`
	row := ur.db.QueryRow(sqlStatement, user.Email)
	err := row.Scan(&id, &hashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Zero rows found")
			return -1, "", err
		} else {
			panic(err)
		}
	}
	return id, hashPassword, nil
}
func (ur *UserRepository) Logout(session string) error {

	sqlStatement := `DELETE FROM sessions WHERE uuid=?;`
	_, err := ur.db.Exec(sqlStatement, session)
	if err != nil {
		return err
	}
	fmt.Println("delete session in db")
	return nil
}
func (ur *UserRepository) UpdateSession(session *models.Session) error {
	//userid, uuid, same UserId -> remove Db, then create New-> row -> session table
	//logout -> remove cookie by userId, cookie delete browser || expires time

	uid := strconv.Itoa(session.UserID)
	_, err := ur.GetUuidInDb(uid)
	if err == nil {
		// return error.New("insert new uuid")
		query, err := ur.db.Prepare(`UPDATE sessions SET uuid=?, cookie_time=? WHERE user_id=?`)
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = query.Exec(session.UUID, session.StartTimeCookie, session.UserID)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		query, err := ur.db.Prepare(`INSERT INTO sessions(uuid, user_id, cookie_time) VALUES(?,?,?)`)
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = query.Exec(session.UUID, session.UserID, session.StartTimeCookie)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (ur *UserRepository) GetUserById(uid string) (*models.User, error) {

	user := models.User{}
	query := `SELECT full_name, email, user_name, age, sex, city FROM users WHERE id=?`
	row := ur.db.QueryRow(query, uid)
	err := row.Scan(&user.FullName, &user.Email, &user.Username, &user.Age, &user.Sex, &user.City)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUuidInDb(uid string) (string, error) {

	var uuid string
	query := `SELECT uuid FROM sessions WHERE user_id=?`
	row := ur.db.QueryRow(query, uid)
	err := row.Scan(&uuid)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

//DRY func ?
func (ur *UserRepository) GetCreatedUserPosts(userId int) (*[]models.Post, error) {

	arrPosts := []models.Post{}
	post := models.Post{}
	var rows *sql.Rows
	var err error
	rows, err = ur.db.Query("SELECT id, thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts WHERE creator_id=?  ORDER  BY create_time  DESC ", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.Thread, &post.Content, &post.CreatorID, &post.CreatedTime, &post.UpdatedTime, &post.Image, &post.CountLike, &post.CountDislike); err != nil {
			return nil, err
		}
		arrPosts = append(arrPosts, post)
	}
	return &arrPosts, nil
}

func (ur *UserRepository) GetUserVotedItems(userId int) (*[]models.Vote, error) {

	seqVote := []models.Vote{}
	voteObj := models.Vote{}
	var rows *sql.Rows
	var err error

	rows, err = ur.db.Query("SELECT post_id, like_state, dislike_state FROM votes WHERE user_id=?", userId)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&voteObj.ID, &voteObj.LikeState, &voteObj.DislikeState); err != nil {
			return nil, err
		}
		if voteObj.LikeState || voteObj.DislikeState {
			seqVote = append(seqVote, voteObj)
		}
	}
	return &seqVote, nil
}
