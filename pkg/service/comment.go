package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/anatolethien/forum/pkg/models"
	"github.com/anatolethien/forum/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo}
}

func (cs *CommentService) Create(comment *models.Comment, postID string) (int, int, error) {
	if err := cs.validateParams(comment); err != nil {
		return 400, -1, err
	}

	postIDint, err := strconv.Atoi(postID)

	comment.CreatedDate = time.Now()
	comment.UpdatedDate = comment.CreatedDate
	comment.Likes = 0
	comment.Dislikes = 0
	comment.PostID = postIDint

	id, err := cs.repo.Create(comment)
	if err != nil {
		return 500, -1, err
	}

	return 200, int(id), nil
}

func (cs *CommentService) EstimateComment(commentID string, userID int, types string) error {
	if types != "like" && types != "dislike" {
		return errors.New("Invalid Type")
	}

	commentIDint, err := strconv.Atoi(commentID)
	if err != nil {
		return err
	}

	comment := &models.Comment{
		ID:     commentIDint,
		UserID: userID,
	}

	return cs.repo.EstimateComment(comment, types)
}

func (cs *CommentService) validateParams(comment *models.Comment) error {
	if comment.Content == "" {
		return errors.New("Invalid Content")
	}

	if comment.UserID < 0 || comment.PostID < 0 {
		return errors.New("Invalid Id")
	}

	return nil
}
