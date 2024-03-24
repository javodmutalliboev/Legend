package model

import "Legend/database"

func (m *Menu) MenuUpdate() error {
	database := database.DB()
	defer database.Close()

	_, err := database.Exec("UPDATE menu SET title = $1, updated_at = NOW() WHERE id = $2", m.Title, m.ID)
	if err != nil {
		return err
	}

	return nil
}
