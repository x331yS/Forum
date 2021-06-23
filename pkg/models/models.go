package models

import (
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	CreatedDate time.Time `json:"created_date"`
}

type Post struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Likes       int        `json:"likes"`
	Dislikes    int        `json:"dislikes"`
	Categories  []string   `json:"categories"`
	Comments    []*Comment `json:"comments"`
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`

	AuthorUsername string
	FormatTime     string
	Images         []string
}

type Comment struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	PostID      int       `json:"post_id"`
	Content     string    `json:"content"`
	Likes       int       `json:"likes"`
	Dislikes    int       `json:"dislikes"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`

	AuthorUsername string
	FormatTime     string
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Session struct {
	UserID  int       `json:"user_id"`
	Token   string    `json:"token"`
	ExpTime time.Time `json:"exp_time"`
}
