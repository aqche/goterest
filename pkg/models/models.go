package models

import "errors"

var ErrUserExists error = errors.New("models: user already exists")

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
