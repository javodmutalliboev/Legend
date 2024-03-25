package model

import "Legend/database"

func (m *Menu) MenuSave() error {
	// Save the menu to the database
	// Open a database connection
	// Insert the menu to the database
	// Close the database connection
	database := database.DB()
	defer database.Close()

	// save title to the database
	_, err := database.Exec("INSERT INTO menu (title, type) VALUES ($1, $2)", m.Title, m.Type)
	if err != nil {
		return err
	}

	return nil
}
