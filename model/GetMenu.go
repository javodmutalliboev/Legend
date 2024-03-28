package model

import "Legend/database"

func GetMenu(id int) (*Menu, error) {
	db := database.DB()
	defer db.Close()

	var menu Menu
	err := db.QueryRow("SELECT * FROM menu WHERE id = $1", id).Scan(&menu.ID, &menu.ParentID, &menu.Title, &menu.CreatedAt, &menu.UpdatedAt, &menu.Type)
	if err != nil {
		return nil, err
	}

	// get children
	err = getChildren(&menu)
	if err != nil {
		return nil, err
	}

	return &menu, nil
}
