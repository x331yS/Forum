package repository

import (
	"database/sql"

	"github.com/anatolethien/forum/pkg/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

func (pr *PostRepository) Create(post *models.Post) (int64, error) {
	tx, err := pr.db.Begin()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	result, err := tx.Exec(`
	INSERT INTO post (user_id,title,content,likes,dislikes,created_date,updated_date)
	VALUES (?,?,?,?,?,?,?)`, post.UserID, post.Title, post.Content, post.Likes, post.Dislikes, post.CreatedDate, post.UpdatedDate)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	for _, category := range post.Categories {
		_, err = tx.Exec(`INSERT INTO category_posts (category_name, post_id) VALUES (?,?)`, category, id)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return id, nil
}

func (pr *PostRepository) GetPostByID(id int) (*models.Post, error) {
	post := &models.Post{}
	if err := pr.db.QueryRow(`
		SELECT post.id, user_id,title,content,likes,dislikes,post.created_date,updated_date,user.username FROM post INNER JOIN user ON user_id=user.id WHERE post.id = ?
	`, id).Scan(&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.Likes,
		&post.Dislikes,
		&post.CreatedDate,
		&post.UpdatedDate,
		&post.AuthorUsername); err != nil {
		return nil, err
	}

	return post, nil
}

func (pr *PostRepository) GetPostsCategories(id int) ([]string, error) {
	categories := []string{}

	rows, err := pr.db.Query("SELECT category_name FROM category_posts WHERE post_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		categories = append(categories, name)
	}

	return categories, nil
}

func (pr *PostRepository) CreateImage(id int, path string) error {
	tx, err := pr.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO post_images (post_id,path)
	VALUES (?,?)`, id, path)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (pr *PostRepository) GetPostsImages(id int) ([]string, error) {
	paths := []string{}

	rows, err := pr.db.Query("SELECT path FROM post_images WHERE post_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		path := ""
		if err := rows.Scan(&path); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return paths, nil
}

func (pr *PostRepository) GetCommentsByPostID(id int) ([]*models.Comment, error) {
	comments := []*models.Comment{}

	rows, err := pr.db.Query("SELECT comment.id,user_id, post_id,content,likes,dislikes,comment.updated_date,user.username FROM comment INNER JOIN user ON user_id=user.id WHERE post_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := &models.Comment{}
		if err := rows.Scan(&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
			&comment.Likes,
			&comment.Dislikes,
			&comment.UpdatedDate,
			&comment.AuthorUsername); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (pr *PostRepository) GetAll() ([]*models.Post, error) {
	posts := []*models.Post{}

	rows, err := pr.db.Query(`
		SELECT post.id, user_id,title,content,likes,dislikes,post.created_date,updated_date,user.username FROM post INNER JOIN user ON user_id=user.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.Dislikes,
			&post.CreatedDate,
			&post.UpdatedDate,
			&post.AuthorUsername); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRepository) GetPostsByCategory(category string) ([]*models.Post, error) {
	posts := []*models.Post{}

	rows, err := pr.db.Query(`
	SELECT post.id, user_id,title,content,likes,dislikes,post.created_date,updated_date,user.username
	FROM post
	INNER JOIN user ON user_id=user.id
	INNER JOIN category_posts ON category_posts.post_id = post.id
	WHERE category_posts.category_name = ?
	`, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.Dislikes,
			&post.CreatedDate,
			&post.UpdatedDate,
			&post.AuthorUsername); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRepository) GetMyCreatedPosts(id int) ([]*models.Post, error) {
	posts := []*models.Post{}

	rows, err := pr.db.Query(`
	SELECT post.id, user_id,title,content,likes,dislikes,post.created_date,updated_date,user.username
	FROM post
	INNER JOIN user ON user_id=user.id
	WHERE user_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.Dislikes,
			&post.CreatedDate,
			&post.UpdatedDate,
			&post.AuthorUsername); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRepository) GetMyLikedPosts(id int) ([]*models.Post, error) {
	posts := []*models.Post{}

	rows, err := pr.db.Query(`
	SELECT post.id,title,content,likes,dislikes,post.created_date,updated_date,user.username
	FROM post
	INNER JOIN post_votes ON post_votes.post_id = post.id
	INNER JOIN user ON user.id = post.user_id
	WHERE type = 'like' AND post_votes.user_id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.Dislikes,
			&post.CreatedDate,
			&post.UpdatedDate,
			&post.AuthorUsername); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (pr *PostRepository) EstimatePost(post *models.Post, types string) error {
	typ := ""

	if err := pr.db.QueryRow("SELECT type FROM post_votes WHERE post_id = ? AND user_id = ?", post.ID, post.UserID).Scan(&typ); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	tx, err := pr.db.Begin()
	if err != nil {
		return err
	}
	if typ == "" {
		_, err = tx.Exec(`
		INSERT INTO post_votes (user_id,post_id,type)
		VALUES (?,?,?)`, post.UserID, post.ID, types)
		if err != nil {
			tx.Rollback()
			return err
		}

		if types == "like" {
			if err := pr.likesChange(tx, post.ID, post.UserID, true); err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err := pr.dislikesChange(tx, post.ID, post.UserID, true); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if typ == types {
		if err := pr.deleteRate(tx, post.ID, post.UserID); err != nil {
			tx.Rollback()
			return err
		}

		if typ == "like" {
			if err := pr.likesChange(tx, post.ID, post.UserID, false); err != nil {
				tx.Rollback()
				return err
			}
		} else {
			if err := pr.dislikesChange(tx, post.ID, post.UserID, false); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if typ == "like" && types == "dislike" {
		if err := pr.likesChange(tx, post.ID, post.UserID, false); err != nil {
			tx.Rollback()
			return err
		}
		if err := pr.dislikesChange(tx, post.ID, post.UserID, true); err != nil {
			tx.Rollback()
			return err
		}
	}

	if typ == "dislike" && types == "like" {
		if err := pr.dislikesChange(tx, post.ID, post.UserID, false); err != nil {
			tx.Rollback()
			return err
		}
		if err := pr.likesChange(tx, post.ID, post.UserID, true); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (pr *PostRepository) GetValidCategories() ([]string, error) {
	categories := []string{}

	rows, err := pr.db.Query("SELECT name FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		categories = append(categories, name)
	}

	return categories, nil
}

func (pr *PostRepository) likesChange(tx *sql.Tx, postID, userID int, up bool) error {
	if up {
		_, err := tx.Exec(`
		UPDATE post SET likes = likes+1 WHERE id = ?`, postID)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(`
		UPDATE post_votes SET type = 'like' WHERE post_id = ? AND user_id`, postID, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err := tx.Exec(`
		UPDATE post SET likes = likes-1 WHERE id = ?`, postID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

func (pr *PostRepository) dislikesChange(tx *sql.Tx, postID, userID int, up bool) error {
	if up {
		_, err := tx.Exec(`
		UPDATE post SET dislikes = dislikes+1 WHERE id = ?`, postID)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(`
		UPDATE post_votes SET type = 'dislike' WHERE post_id = ? AND user_id`, postID, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err := tx.Exec(`
		UPDATE post SET dislikes = dislikes-1 WHERE id = ?`, postID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

func (pr *PostRepository) deleteRate(tx *sql.Tx, postID, userID int) error {
	_, err := tx.Exec(`
	DELETE FROM post_votes WHERE user_id = ? AND post_id = ?`, userID, postID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
