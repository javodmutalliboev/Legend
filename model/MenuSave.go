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
	_, err := database.Exec("INSERT INTO menu (title_uz, title_ru, title_en, type) VALUES ($1, $2, $3, $4)", m.TitleUz, m.TitleRu, m.TitleEn, m.Type)
	if err != nil {
		return err
	}

	return nil
}
