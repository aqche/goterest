package main

import (
	"errors"
	"os"
)

var ErrMissingDatabaseURL error = errors.New("Missing DATABASE_URL environment variable.")
var ErrMissingSessionKey error = errors.New("Missing SESSION_KEY environment variable.")
var ErrMissingCSRFKey error = errors.New("Missing CSRF_KEY environment variable.")

type config struct {
	Port        string
	DatabaseURL string
	SessionKey  string
	CSRFKey     string
	Secure      bool
}

func getConfig() (*config, error) {
	config := &config{}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "4000"
	}

	config.DatabaseURL = os.Getenv("DATABASE_URL")
	if config.DatabaseURL == "" {
		return nil, ErrMissingDatabaseURL
	}

	config.SessionKey = os.Getenv("SESSION_KEY")
	if config.SessionKey == "" {
		return nil, ErrMissingSessionKey
	}

	config.CSRFKey = os.Getenv("CSRF_KEY")
	if config.CSRFKey == "" {
		return nil, ErrMissingCSRFKey
	}

	if os.Getenv("ENV") == "DEV" {
		config.Secure = false
	} else {
		config.Secure = true
	}

	return config, nil
}
