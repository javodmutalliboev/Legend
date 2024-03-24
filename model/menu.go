package model

import "Legend/database"

type Menu struct {
	ID        int    `json:"id"`
	ParentID  int    `json:"parent_id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (m *Menu) MenuSave() error {
	// Save the menu to the database
	// Open a database connection
	// Insert the menu to the database
	// Close the database connection
	database := database.DB()
	defer database.Close()

	// save title to the database
	_, err := database.Exec("INSERT INTO menu (title) VALUES ($1)", m.Title)
	if err != nil {
		return err
	}

	return nil
}
