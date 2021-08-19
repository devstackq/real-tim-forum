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
