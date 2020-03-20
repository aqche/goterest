package postgres

import "database/sql"

type UserModel struct {
	DB *sql.DB
}
