package model

import (
	"Legend/database"
	"errors"
)

type CTWInformation struct {
	ID          int    `json:"id"`
	Heading     string `json:"heading"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CreateCTWInformation(ctw *CTWInformation) error {
	db := database.DB()
	defer db.Close()

	// check for at least one ctw_information row exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM ctw_information)").Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// return error of 'ctw information already exists'
		return errors.New("ctw information already exists")
	}

	_, err = db.Exec("INSERT INTO ctw_information (heading, description) VALUES ($1, $2)", ctw.Heading, ctw.Description)
	if err != nil {
		return err
	}

	return nil
}
