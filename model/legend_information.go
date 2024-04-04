package model

import (
	"Legend/database"
	"errors"
)

type LegendInformation struct {
	ID          int    `json:"id"`
	Heading     string `json:"heading"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CreateLegendInformation(li *LegendInformation) error {
	db := database.DB()
	defer db.Close()

	// check for at least one legend_information row exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM legend_information)").Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// return error of 'legend information already exists'
		return errors.New("legend information already exists")
	}

	_, err = db.Exec("INSERT INTO legend_information (heading, description) VALUES ($1, $2)", li.Heading, li.Description)
	if err != nil {
		return err
	}

	return nil
}

func GetLegendInformation() (LegendInformation, error) {
	db := database.DB()
	defer db.Close()

	var li LegendInformation
	err := db.QueryRow("SELECT * FROM legend_information").Scan(&li.ID, &li.Heading, &li.Description, &li.CreatedAt, &li.UpdatedAt)
	if err != nil {
		return li, err
	}

	return li, nil
}
