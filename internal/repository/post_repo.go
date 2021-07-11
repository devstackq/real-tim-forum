package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

//have db - connect
type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

func (ur *PostRepository) GetUserID(uuid string) (int, error) {
	var userid int
	query := `SELECT user_id FROM sessions WHERE uuid=?`
	row := ur.db.QueryRow(query, uuid)
	err := row.Scan(&userid)
	if err != nil {
		return 0, err
	}
	return userid, nil
}

func (pr *PostRepository) CreatePost(post *models.Post) (string, int, error) {
	query, err := pr.db.Prepare(`INSERT INTO posts(thread, content, creator_id) VALUES(?,?,?)`)
	if err != nil {
		return "", -1, err
	}
	result, err := query.Exec(
		post.Thread,
		post.Content,
		post.CreatorID,
	)
	// post.image
	if err != nil {
		return "", -1, err
	}

	defer query.Close()

	postid, err := result.LastInsertId()
	post.ID = fmt.Sprint(postid)
	if err != nil {
		return "", -1, err
	}
	// log.Printf("Created a new post with id %d\n", postid)
	return fmt.Sprint(postid), http.StatusOK, nil
}

func (pr *PostRepository) JoinCategoryPost(pid string, category string) error {
	// fmt.Println(pid, category)
	query, err := pr.db.Prepare(`INSERT INTO post_category_bridges(post_id, category_id) VALUES(?,?)`)
	if err != nil {
		return err
	}
	_, err = query.Exec(pid, category)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRepository) GetPostsByCategory(category string) (*[]models.Post, error) {

	arrPosts := []models.Post{}

	post := models.Post{}
	var rows *sql.Rows
	var err error
	//create post -> save now ->posts.cat_id 1/2/3/ etc
	// get all posts -> select *  posts.cat-id = categories.id, post FK(cat_id) Ref(categories(id))

	if category == "love" {
		rows, err = pr.db.Query("SELECT posts.id, thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts LEFT JOIN post_category_bridges  ON post_category_bridges.post_id = posts.id  WHERE category_id=?  ORDER  BY create_time  DESC", 2)
	} else if category == "science" {
		rows, err = pr.db.Query("SELECT  posts.id, thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts LEFT JOIN post_category_bridges  ON post_category_bridges.post_id = posts.id  WHERE category_id=?  ORDER  BY create_time  DESC", 1)
	} else if category == "nature" {
		rows, err = pr.db.Query("SELECT posts.id, thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts LEFT JOIN post_category_bridges  ON post_category_bridges.post_id = posts.id WHERE category_id=?  ORDER  BY create_time  DESC", 3)
	} else if category == "all" || category == "" {
		rows, err = pr.db.Query("SELECT id, thread, content, creator_id, create_time, update_time, image, count_like, count_dislike FROM posts ORDER  BY create_time  DESC")
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		//query getUsername byPostId
		if err := rows.Scan(&post.ID, &post.Thread, &post.Content, &post.CreatorID, &post.CreatedTime, &post.UpdatedTime, &post.Image, &post.CountLike, &post.CountDislike); err != nil {
			return nil, err
		}
		arrPosts = append(arrPosts, post)
	}
	return &arrPosts, nil
}

func (pr *PostRepository) GetPostById(postId string) (*models.Post, error) {

	post := models.Post{}
	query := `SELECT * FROM posts WHERE id=?`
	row := pr.db.QueryRow(query, postId)
	err := row.Scan(&post.ID, &post.Thread, &post.Content, &post.CreatorID, &post.CreatedTime, &post.UpdatedTime, &post.Image, &post.CountLike, &post.CountDislike)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
