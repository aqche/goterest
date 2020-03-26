package models

import "errors"

var (
	ErrUserExists   error = errors.New("models: user already exists")
	ErrUserNotFound error = errors.New("models: user not found")
)

type User struct {
	ID       int
	Username string
	Password []byte
}

type Pin struct {
	ID       int
	ImageURL string
	Username string
}
