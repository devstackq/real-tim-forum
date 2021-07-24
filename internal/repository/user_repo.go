package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

//have db - connect
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	fmt.Println("create user repo ")
	return &UserRepository{db}
}

//implement method, by interface User
func (ur *UserRepository) CreateUser(user *models.User) (int64, error) {

	query, err := ur.db.Prepare(`INSERT INTO users(full_name, email, user_name, password, age, sex, created_time, last_seen, city, image) VALUES(?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return -1, err
	}
	t := time.Now().Format(time.RFC822)

	result, err := query.Exec(user.FullName, user.Email, user.Username, user.Password, user.Age, user.Sex, t, time.Now(), user.City, user.Image)
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
	_, err := ur.GetUserUUID(session.UserID)
	// delete in global variable

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
	query := `SELECT full_name, email, user_name, age, sex, created_time, city FROM users WHERE id=?`
	row := ur.db.QueryRow(query, uid)
	err := row.Scan(&user.FullName, &user.Email, &user.Username, &user.Age, &user.Sex, &user.CreatedTime, &user.City)
	if err != nil {
		return nil, err
	}
	log.Println(user.CreatedTime, "time")
	return &user, nil
}

func (ur *UserRepository) GetUserUUID(uid int) (string, error) {
	var uuid string
	query := `SELECT uuid FROM sessions WHERE user_id=?`
	row := ur.db.QueryRow(query, uid)
	err := row.Scan(&uuid)

	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (ur *UserRepository) GetUserName(id string) (string, error) {

	var name string
	query := `SELECT full_name FROM users WHERE id=?`
	row := ur.db.QueryRow(query, id)
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

//dry?
func (ur *UserRepository) GetUserID(uuid string) (int, error) {
	var userid int
	query := `SELECT user_id FROM session WHERE uuid=?`
	row := ur.db.QueryRow(query, uuid)
	err := row.Scan(&userid)
	if err != nil {
		return 0, err
	}
	return userid, nil
}

func (ur *UserRepository) GetProfileData(userId int) (*models.Profile, error) {
	//by user id , get created posts, vote, and user data
	//2 query, 1 get all post
	queryStmt, err := ur.db.Query(`SELECT *  FROM users u
	left join posts p ON p.user_id = $1
	left join votes v ON v.user_id = $1  ORDER by p.create_time DESC`, userId)
	if err != nil {
		return nil, err
	}
	//get user data
	// SELECT * FROM users where user_id=?
	//get voted posts
	// SELECT u.full_name, v.post_id as postid, u.email, v.id as voteid, u.id as userid  FROM users u
	// left join votes v ON v.user_id = u.id WHERE
	// u.id =5   ORDER by v.id DESC

	//created post
	// SELECT u.full_name, u.email, p.thread, u.id as userid, p.id as postid  FROM users u
	// left join posts p ON p.creator_id = u.id   where u.id=5

	// SELECT u.full_name, p.thread, v.id as voteid, u.id as userid, p.id as postid  FROM users u
	// left join posts p ON p.creator_id = u.id
	// left join votes v ON v.user_id = u.id AND v.post_id NOTNULL  WHERE u.id =5 GROUP BY u.id, v.id, p.id  ORDER by p.create_time DESC

	// profile := []*models.Profile{}
	// for queryStmt.Next() {
	// 	u := models.User{}
	// 	if err := queryStmt.Scan(&c.ID, &c.UserName, &c.LastMessage, &c.LastMessageSenderName, &c.SentTime, &c.MesageID); err != nil {
	// 		return nil, err
	// 	}
	// 	chatUsers = append(chatUsers, &c)
	// }
	// return chatUsers, nil
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
