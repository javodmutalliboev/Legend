package model

import (
	"Legend/database"
	"errors"
	"fmt"
	"strings"
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

func UpdateLegendInformation(li *LegendInformation) error {
	db := database.DB()
	defer db.Close()

	var fields []string
	var args []interface{}
	i := 1

	if li.Heading != "" {
		fields = append(fields, fmt.Sprintf("heading = $%d", i))
		args = append(args, li.Heading)
		i++
	}

	if li.Description != "" {
		fields = append(fields, fmt.Sprintf("description = $%d", i))
		args = append(args, li.Description)
		i++
	}

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	// build the SQL query
	sql := fmt.Sprintf("UPDATE legend_information SET %s, updated_at = NOW() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, li.ID)

	_, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
