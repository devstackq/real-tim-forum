package repository

import (
	"database/sql"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db}
}

//get pid, insert data, retun inserted conetent, idcomment, id post
func (cr *CommentRepository) CreateComment(c *models.Comment) (*models.Comment, error) {
	query, err := cr.db.Prepare(`INSERT INTO comments(post_id, content, creator_id, create_time) VALUES(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	c.CreatedTime = time.Now()

	id, err := query.Exec(
		c.PostID,
		c.Content,
		c.CreatorID,
		c.CreatedTime,
	)
	if err != nil {
		return nil, err
	}
	lid, _ := id.LastInsertId()
	c.ID = int(lid)
	return c, nil
}

func (cr *CommentRepository) GetCommentsByID(pid int) (*[]models.Comment, error) {

	comment := models.Comment{}
	seqComments := []models.Comment{}
	rows, err := cr.db.Query("SELECT id,  content, post_id, creator_id, create_time, count_like, count_dislike FROM comments WHERE post_id=?", pid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&comment.ID, &comment.Content, &comment.PostID, &comment.CreatorID, &comment.CreatedTime, &comment.CountLike, &comment.CountDislike); err != nil {
			return nil, err
		}
		seqComments = append(seqComments, comment)
	}
	return &seqComments, nil
}
