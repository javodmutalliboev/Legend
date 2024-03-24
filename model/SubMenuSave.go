package model

import "Legend/database"

func (m *Menu) SubMenuSave() error {
	database := database.DB()
	defer database.Close()

	// save parent_id and title to the database
	_, err := database.Exec("INSERT INTO menu (parent_id, title) VALUES ($1, $2)", m.ParentID, m.Title)
	if err != nil {
		return err
	}

	return nil
}
