package repository

import (
	"database/sql"

	"github.com/anatolethien/forum/pkg/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *models.User) (int64, error) {
	result, err := ur.db.Exec(`
	INSERT INTO user (email,username,password,role,created_date)
	VALUES (?,?,?,?,?)`, user.Email, user.Username, user.Password, user.Role, user.CreatedDate)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.QueryRow(`
		SELECT id,username,email,role,password,created_date FROM user WHERE email = ?
	`, email).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Password, &user.CreatedDate); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.QueryRow(`
		SELECT id,username,email,role,password,created_date FROM user WHERE username = ?
	`, username).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Password, &user.CreatedDate); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.QueryRow(`
		SELECT id,username,email,role,password,created_date FROM user WHERE id = ?
	`, id).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Password, &user.CreatedDate); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) CreateSession(session *models.Session) error {
	if _, err := ur.db.Exec(`
	INSERT INTO session (user_id,token,exp_time)
	VALUES (?,?,?)`, session.UserID, session.Token, session.ExpTime); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdateSession(session *models.Session) error {
	if _, err := ur.db.Exec(`
		UPDATE session SET token = ?, exp_time = ?, user_id = ? WHERE user_id = ?`,
		session.Token, session.ExpTime, session.UserID, session.UserID); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteSession(token string) error {
	if _, err := ur.db.Exec(`
		DELETE FROM session WHERE token = ?`, token); err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetSession(token string) (*models.Session, error) {
	session := &models.Session{}
	if err := ur.db.QueryRow(`
		SELECT user_id,token,exp_time FROM session WHERE token = ?
	`, token).Scan(&session.UserID, &session.Token, &session.ExpTime); err != nil {
		return nil, err
	}
	return session, nil
}
