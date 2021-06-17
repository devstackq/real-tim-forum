package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

//have db - connect
type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

//implement method, by interface Post
func (pr *PostRepository) CreatePost(post *models.Post) (int, error) {

	// log.Printf("Creating new post for userid %d...\n", post.CreatorID)
	fmt.Println(post)
	query, err := pr.db.Prepare(`
		INSERT INTO posts(
			thread, content, creator_id, create_time, image, count_like, count_dislike
		) VALUES(?,?,?,?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return -1, err
	}

	result, err := query.Exec(
		post.Thread,
		post.Content,
		post.CreatorID,
		time.Now(),
		post.Image,
		post.CountLike,
		post.CountDislike,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	defer query.Close()

	postid, err := result.LastInsertId()
	post.ID = int(postid)
	if err != nil {
		return -1, err
	}

	err = pr.JoinCategoryPost(postid, post.Category)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// or redirect -> by pid ?
	log.Printf("Created a new post with id %d\n", postid)

	return http.StatusOK, nil
}

func (pr *PostRepository) JoinCategoryPost(pid int64, category string) error {

	query, err := pr.db.Prepare(`INSERT INTO post_category_bridge(post_id, category_id) VALUES(?,?)`)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = query.Exec(pid, category)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pr *PostRepository) GetPostsByCategory(category string) (*[]models.Post, error) {

	arrPosts := []models.Post{}
	post := models.Post{}
	var rows *sql.Rows
	var err error

	if category == "/" {
		rows, err = pr.db.Query("SELECT posts.id , thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts ORDER  BY create_time  DESC")
	} else if category == "love" {
		rows, err = pr.db.Query("SELECT posts.id , thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts LEFT JOIN post_category_bridge  ON post_category_bridge.post_id = posts.id   WHERE category_id=?  ORDER  BY create_time  DESC", 2)
	} else if category == "science" {
		rows, err = pr.db.Query("SELECT posts.id , thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts LEFT JOIN post_category_bridge  ON post_category_bridge.post_id = posts.id   WHERE category_id=?  ORDER  BY create_time  DESC", 1)
	} else if category == "nature" {
		rows, err = pr.db.Query("SELECT posts.id , thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts LEFT JOIN post_category_bridge  ON post_category_bridge.post_id = posts.id   WHERE category_id=?  ORDER  BY create_time  DESC", 3)
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		if err := rows.Scan(&post.ID, &post.Thread, &post.Content, &post.CreatorID, &post.CreatedTime, &post.UpdatedTime, &post.Image, &post.CountLike, &post.CountDislike); err != nil {
			return nil, err
		}
		arrPosts = append(arrPosts, post)
	}
	// fmt.Println(arrPosts)
	return &arrPosts, nil
}
