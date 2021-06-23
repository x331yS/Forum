package repository

import (
	"database/sql"

	"github.com/anatolethien/forum/pkg/models"
)

type User interface {
	CreateUser(*models.User) (int64, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserByUsername(string) (*models.User, error)
	GetUserByID(int) (*models.User, error)

	CreateSession(*models.Session) error
	UpdateSession(*models.Session) error
	DeleteSession(string) error
	GetSession(string) (*models.Session, error)
}

type Post interface {
	Create(*models.Post) (int64, error)
	GetAll() ([]*models.Post, error)
	GetPostByID(int) (*models.Post, error)
	GetValidCategories() ([]string, error)
	GetPostsCategories(int) ([]string, error)
	GetMyCreatedPosts(int) ([]*models.Post, error)
	GetMyLikedPosts(int) ([]*models.Post, error)
	GetPostsByCategory(string) ([]*models.Post, error)
	EstimatePost(*models.Post, string) error
	CreateImage(int, string) error
	GetPostsImages(int) ([]string, error)

	GetCommentsByPostID(int) ([]*models.Comment, error)
}

type Comment interface {
	Create(*models.Comment) (int64, error)
	EstimateComment(*models.Comment, string) error
}

type Repository struct {
	User
	Post
	Comment
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Post:    NewPostRepository(db),
		Comment: NewCommentRepository(db),
	}
}
