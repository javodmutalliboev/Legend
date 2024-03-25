package model

import (
	"Legend/database"
	"errors"
)

func (m *Menu) SubMenuSave() error {
	database := database.DB()
	defer database.Close()

	// check if parent menu type is same as the current menu type
	var parentType int
	err := database.QueryRow("SELECT type FROM menu WHERE id = $1", m.ParentID).Scan(&parentType)
	if err != nil {
		return err
	}

	if parentType != m.Type {
		return errors.New("sub menu type must be the same as the parent menu type")
	}

	// save parent_id and title to the database
	_, err = database.Exec("INSERT INTO menu (parent_id, title, type) VALUES ($1, $2, $3)", m.ParentID, m.Title, m.Type)
	if err != nil {
		return err
	}

	return nil
}
