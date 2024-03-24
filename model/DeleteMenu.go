package model

import "Legend/database"

func DeleteMenu(id int) error {
	db := database.DB()
	defer db.Close()

	// first remove dependents: reference column is parent_id
	_, err := db.Exec("DELETE FROM menu WHERE parent_id = $1", id)
	if err != nil {
		return err
	}

	// then remove the menu
	_, err = db.Exec("DELETE FROM menu WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
