package repository

import (
	"database/sql"
	"log"
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
func (ur *PostRepository) CreatePost(post *models.Post) (int64, error) {

	log.Printf("Creating new post for userid %d...\n", post.CreatorID)

	query, err := ur.db.Prepare(`
		INSERT INTO posts(
			thread, content, create_time, update_time, image, count_like, count_dislike
		) VALUES(?,?,?,?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return -1, err
	}

	result, err := query.Exec(
		post.Thread,
		post.Content,
		time.Now(),
		post.UpdatedTime,
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

	// Add to categories
	// err = insertPostIntoCategories(post.ID, post.Categories)
	// if err != nil {
	// 	return -1, err
	// }

	log.Printf("Created a new post with id %d\n", postid)

	return result.LastInsertId()
}
