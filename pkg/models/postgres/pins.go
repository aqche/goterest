package postgres

import (
	"database/sql"

	"github.com/aqche/goterest/pkg/models"
)

type PinModel struct {
	DB *sql.DB
}

func (m *PinModel) Create(imageURL string, username string) error {
	stmt := "INSERT INTO pins (image_url, username) VALUES ($1, $2)"

	_, err := m.DB.Exec(stmt, imageURL, username)
	if err != nil {
		return err
	}

	return nil
}

func (m *PinModel) GetAll() ([]*models.Pin, error) {
	stmt := "SELECT pin_id, image_url, username FROM pins"

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pins := []*models.Pin{}

	for rows.Next() {
		pin := &models.Pin{}

		err := rows.Scan(&pin.ID, &pin.ImageURL, &pin.Username)
		if err != nil {
			return nil, err
		}

		pins = append(pins, pin)
	}

	return pins, nil
}

func (m *PinModel) GetAllByUsername(username string) ([]*models.Pin, error) {
	stmt := "SELECT pin_id, image_url, username FROM pins WHERE username = $1"

	rows, err := m.DB.Query(stmt, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pins := []*models.Pin{}

	for rows.Next() {
		pin := &models.Pin{}

		err := rows.Scan(&pin.ID, &pin.ImageURL, &pin.Username)
		if err != nil {
			return nil, err
		}

		pins = append(pins, pin)
	}

	return pins, nil
}
