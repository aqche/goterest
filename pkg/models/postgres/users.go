package postgres

import (
	"database/sql"
	"errors"

	"github.com/aqche/goterest/pkg/models"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Create(username string, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := "INSERT INTO users (username, password) VALUES ($1, $2)"

	_, err = m.DB.Exec(stmt, username, passwordHash)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return models.ErrUserExists
			}
		}
		return err
	}

	return nil
}

func (m *UserModel) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}

	stmt := "SELECT user_id, username, password FROM users WHERE username = $1"

	row := m.DB.QueryRow(stmt, username)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (m *UserModel) ValidatePassword(username string, password string) error {
	user, err := m.GetByUsername(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.ErrInvalidPassword
		}
		return err
	}

	return nil
}
