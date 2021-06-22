package repository

import (
	"database/sql"

	"github.com/anatolethien/forum/internal/app/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (cr *CommentRepository) Create(comment *models.Comment) (int64, error) {
	tx, err := cr.db.Begin()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	result, err := tx.Exec(`
	INSERT INTO comment (user_id,post_id,content,likes,dislikes,created_date,updated_date) 
	VALUES (?,?,?,?,?,?,?)`,
		comment.UserID, comment.PostID, comment.Content, comment.Likes, comment.Dislikes, comment.CreatedDate, comment.UpdatedDate)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return id, nil
}

func (cr *CommentRepository) EstimateComment(comment *models.Comment, types string) error {
	typ := ""

	if err := cr.db.QueryRow("SELECT type FROM comment_votes WHERE comment_id = ? AND user_id = ?", comment.ID, comment.UserID).Scan(&typ); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	tx, err := cr.db.Begin()
	if err != nil {
		return err
	}

	if typ == "" {
		_, err = tx.Exec(`
		INSERT INTO comment_votes (user_id,comment_id,type) 
		VALUES (?,?,?)`, comment.UserID, comment.ID, types)
		if err != nil {
			tx.Rollback()
			return err
		}

		if types == "like" {
			if err := cr.likesChange(tx, comment.ID, comment.UserID, true); err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err := cr.dislikesChange(tx, comment.ID, comment.UserID, true); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if typ == types {
		if err := cr.deleteRate(tx, comment.ID, comment.UserID); err != nil {
			tx.Rollback()
			return err
		}

		if typ == "like" {
			if err := cr.likesChange(tx, comment.ID, comment.UserID, false); err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err := cr.dislikesChange(tx, comment.ID, comment.UserID, false); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if typ == "like" && types == "dislike" {
		if err := cr.likesChange(tx, comment.ID, comment.UserID, false); err != nil {
			tx.Rollback()
			return err
		}
		if err := cr.dislikesChange(tx, comment.ID, comment.UserID, true); err != nil {
			tx.Rollback()
			return err
		}
	}

	if typ == "dislike" && types == "like" {
		if err := cr.dislikesChange(tx, comment.ID, comment.UserID, false); err != nil {
			tx.Rollback()
			return err
		}
		if err := cr.likesChange(tx, comment.ID, comment.UserID, true); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (cr *CommentRepository) likesChange(tx *sql.Tx, commentID, userID int, up bool) error {
	if up {
		_, err := tx.Exec(`
		UPDATE comment SET likes = likes+1 WHERE id = ?`, commentID)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(`
		UPDATE comment_votes SET type = 'like' WHERE comment_id = ? AND user_id`, commentID, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err := tx.Exec(`
		UPDATE comment SET likes = likes-1 WHERE id = ?`, commentID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

func (cr *CommentRepository) dislikesChange(tx *sql.Tx, commentID, userID int, up bool) error {
	if up {
		_, err := tx.Exec(`
		UPDATE comment SET dislikes = dislikes+1 WHERE id = ?`, commentID)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(`
		UPDATE comment_votes SET type = 'dislike' WHERE comment_id = ? AND user_id`, commentID, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err := tx.Exec(`
		UPDATE comment SET dislikes = dislikes-1 WHERE id = ?`, commentID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

func (cr *CommentRepository) deleteRate(tx *sql.Tx, commentID, userID int) error {
	_, err := tx.Exec(`
	DELETE FROM comment_votes WHERE user_id = ? AND comment_id = ?`, userID, commentID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
