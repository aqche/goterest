package postgres

import "database/sql"

type PinModel struct {
	DB *sql.DB
}
