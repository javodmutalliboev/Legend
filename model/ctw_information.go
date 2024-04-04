package model

import (
	"Legend/database"
	"errors"
	"fmt"
	"strings"
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

func GetCTWInformation() (CTWInformation, error) {
	db := database.DB()
	defer db.Close()

	var ctw CTWInformation
	err := db.QueryRow("SELECT * FROM ctw_information").Scan(&ctw.ID, &ctw.Heading, &ctw.Description, &ctw.CreatedAt, &ctw.UpdatedAt)
	if err != nil {
		return ctw, err
	}

	return ctw, nil
}

func UpdateCTWInformation(ctw *CTWInformation) error {
	db := database.DB()
	defer db.Close()

	var fields []string
	var args []interface{}
	i := 1

	if ctw.Heading != "" {
		fields = append(fields, fmt.Sprintf("heading = $%d", i))
		args = append(args, ctw.Heading)
		i++
	}

	if ctw.Description != "" {
		fields = append(fields, fmt.Sprintf("description = $%d", i))
		args = append(args, ctw.Description)
		i++
	}

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	// build the SQL query
	sql := fmt.Sprintf("UPDATE ctw_information SET %s, updated_at = NOW() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, ctw.ID)

	_, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
