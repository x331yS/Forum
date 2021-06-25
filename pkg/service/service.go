package service

import (
	"mime/multipart"

	"github.com/anatolethien/forum/pkg/models"
	"github.com/anatolethien/forum/pkg/repository"
)

type User interface {
	Create(*models.User) (int, int, error)
	Authorization(string, string) (*models.Session, error)
	Logout(string) error
	IsValidToken(string) bool
	GetUserIDByToken(string) (int, error)
}

type Post interface {
	Create(*models.Post) (int, int, error)
	Get(int) (*models.Post, error)
	GetAll() ([]*models.Post, error)
	GetValidCategories() ([]string, error)
	GetCommentsByPostID(int) ([]*models.Comment, error)
	Filter(string, int) ([]*models.Post, error)
	EstimatePost(string, int, string) error
	SetImage(int, string) error
	GenerateImagesFromFiles([]*multipart.FileHeader) ([]string, error)
}

type Comment interface {
	Create(*models.Comment, string) (int, int, error)
	EstimateComment(string, int, string) error
}

type Service struct {
	User
	Post
	Comment
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(r.User),
		Post:    NewPostService(r.Post),
		Comment: NewCommentService(r.Comment),
	}
}
