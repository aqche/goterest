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

func (m *UserModel) ValidatePassword(username string, password string) error {
	var passwordHash []byte

	stmt := "SELECT password FROM users WHERE username = $1"

	err := m.DB.QueryRow(stmt, username).Scan(&passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrUserNotFound
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.ErrInvalidPassword
		}
		return err
	}

	return nil
}
