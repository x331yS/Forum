package service

import (
	"database/sql"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/anatolethien/forum/internal/app/models"
	"github.com/anatolethien/forum/internal/app/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo}
}

func (ps *PostService) Create(post *models.Post) (int, int, error) {
	if err := ps.validateParams(post); err != nil {
		return 400, -1, err
	}

	post.CreatedDate = time.Now()
	post.UpdatedDate = post.CreatedDate
	post.Likes = 0
	post.Dislikes = 0

	id, err := ps.repo.Create(post)
	if err != nil {
		return 500, -1, err
	}

	return 200, int(id), nil
}

func (ps *PostService) GenerateImagesFromFiles(files []*multipart.FileHeader) ([]string, error) {
	paths := []string{}

	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			return nil, err
		}

		defer file.Close()
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			return nil, err
		}

		fileType := http.DetectContentType(buff)
		if fileType != "image/jpeg" && fileType != "image/png" && fileType != "image/jpg" && fileType != "image/gif" {
			return nil, errors.New("Invalid File Type")
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return nil, errors.New("Invalid File Type")
		}

		err = os.MkdirAll("./assets/images", os.ModePerm)
		if err != nil {
			return nil, err
		}
		imageName := uuid.NewV4().String()
		destImage := fmt.Sprintf("/images/%d%s", imageName, filepath.Ext(files[i].Filename))
		dst, err := os.Create("./assets" + destImage)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			return nil, err
		}

		paths = append(paths, destImage)
	}

	return paths, nil
}

func (ps *PostService) Get(id int) (*models.Post, error) {
	if id < 0 {
		return nil, errors.New("Invalid Id")
	}
	post, err := ps.repo.GetPostByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid Id")
		}
		return nil, err
	}

	postCategories, err := ps.repo.GetPostsCategories(id)
	if err != nil {
		return nil, err
	}

	post.Categories = postCategories

	postComments, err := ps.repo.GetCommentsByPostID(id)
	if err != nil {
		return nil, err
	}

	for i := range postComments {
		postComments[i].FormatTime = postComments[i].UpdatedDate.Format("January 02, 2006")
	}

	postImages, err := ps.repo.GetPostsImages(id)
	if err != nil {
		return nil, err
	}

	post.Images = postImages
	post.Comments = postComments
	post.FormatTime = post.UpdatedDate.Format("January 02, 2006")

	return post, nil
}

func (ps *PostService) GetAll() ([]*models.Post, error) {
	posts, err := ps.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for i, post := range posts {
		post.FormatTime = post.UpdatedDate.Format("January 02, 2006")
		post.Categories, err = ps.repo.GetPostsCategories(post.ID)
		if err != nil {
			return nil, err
		}
		posts[i] = post
	}

	return posts, nil
}

func (ps *PostService) Filter(field string, id int) ([]*models.Post, error) {
	posts := []*models.Post{}
	var err error

	if id == 0 && (field == "Myliked" || field == "Mycreated") {
		return nil, errors.New("Unauthorized")
	}

	if field == "Myliked" {
		posts, err = ps.repo.GetMyLikedPosts(id)
	} else if field == "Mycreated" {
		posts, err = ps.repo.GetMyCreatedPosts(id)
	} else {
		categories, err := ps.repo.GetValidCategories()
		if err != nil {
			return nil, err
		}

		field = strings.Title(strings.ToLower(field))

		ok := false
		for _, c := range categories {
			if field == c {
				ok = true
			}
		}
		if !ok {
			return nil, errors.New("Invalid Category")
		}

		posts, err = ps.repo.GetPostsByCategory(field)
	}

	if err != nil {
		return nil, err
	}

	for i, post := range posts {
		post.FormatTime = post.UpdatedDate.Format("January 02, 2006")
		post.Categories, err = ps.repo.GetPostsCategories(post.ID)
		if err != nil {
			return nil, err
		}
		posts[i] = post
	}

	return posts, nil
}

func (ps *PostService) GetValidCategories() ([]string, error) {
	return ps.repo.GetValidCategories()
}

func (ps *PostService) SetImage(id int, path string) error {
	return ps.repo.CreateImage(id, path)
}

func (ps *PostService) GetCommentsByPostID(id int) ([]*models.Comment, error) {
	comments, err := ps.repo.GetCommentsByPostID(id)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (ps *PostService) EstimatePost(postID string, userID int, types string) error {
	if types != "like" && types != "dislike" {
		return errors.New("Invalid Type")
	}

	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		return err
	}

	post := &models.Post{
		ID:     postIDint,
		UserID: userID,
	}

	return ps.repo.EstimatePost(post, types)
}

func (ps *PostService) validateParams(post *models.Post) error {
	if post.Title == "" || strings.TrimSpace(post.Title) == "" {
		return errors.New("Invalid Title")
	}

	if post.Content == "" || strings.TrimSpace(post.Content) == "" {
		return errors.New("Invalid Content")
	}

	if post.Categories == nil {
		return errors.New("Invalid Category")
	}

	if categories, err := ps.repo.GetValidCategories(); err != nil {
		return err
	} else {
		for _, postCategory := range post.Categories {
			ok := 1
			for _, category := range categories {
				if postCategory == category {
					ok--
				}
			}
			if ok != 0 {
				return errors.New("Invalid Category")
			}
		}
	}

	return nil
}
