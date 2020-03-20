package models

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
