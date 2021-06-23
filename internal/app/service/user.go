package service

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/anatolethien/forum/internal/app/models"
	"github.com/anatolethien/forum/internal/app/repository"
	"golang.org/x/crypto/bcrypt"

	sqlite "github.com/mattn/go-sqlite3"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo}
}

func (us *UserService) Create(user *models.User) (int, int, error) {
	if err := us.validateParams(user); err != nil {
		return http.StatusBadRequest, -1, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, -1, err
	}

	user.Password = string(hashPassword)
	user.CreatedDate = time.Now()

	id, err := us.repo.CreateUser(user)
	if err != nil {
		if sqliteErr, ok := err.(sqlite.Error); ok {
			if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
				return http.StatusBadRequest, -1, errors.New("User already created")
			}
		}
		return http.StatusInternalServerError, -1, err
	}

	return http.StatusOK, int(id), nil
}

func (us *UserService) Authorization(login, password string) (*models.Session, error) {
	user := &models.User{}
	var err error

	if strings.Contains(login, "@") {
		user, err = us.repo.GetUserByEmail(login)
	} else {
		user, err = us.repo.GetUserByUsername(login)
	}

	if err != nil {
		return nil, errors.New("Invalid email/login or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid email/login or password")
	}

	session := &models.Session{
		UserID:  user.ID,
		ExpTime: time.Now().Add(time.Minute * 30),
		Token:   uuid.NewV4().String(),
	}

	if err := us.repo.CreateSession(session); err != nil {
		if sqliteErr, ok := err.(sqlite.Error); ok {
			if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
				if err2 := us.repo.UpdateSession(session); err2 != nil {
					return nil, err2
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return session, nil
}

func (us *UserService) Logout(token string) error {
	return us.repo.DeleteSession(token)
}

func (us *UserService) IsValidToken(token string) bool {
	if s, err := us.repo.GetSession(token); err != nil || s == nil {
		return false
	}
	return true
}

func (us *UserService) GetUserIDByToken(token string) (int, error) {
	if s, err := us.repo.GetSession(token); err != nil {
		return -1, err
	} else {
		return s.UserID, nil
	}
}

func (us *UserService) validateParams(user *models.User) error {
	patternForEmail := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`

	ok, _ := regexp.MatchString(patternForEmail, user.Email)
	if !ok {
		return errors.New("Invalid Email")
	}

	if user.Username == "" || len(user.Username) < 2 {
		return errors.New("Invalid Username")
	}

	if len(user.Password) < 6 {
		return errors.New("Invalid Password")
	}

	if user.Role != "user" {
		return errors.New("Invalid Role")
	}

	return nil
}