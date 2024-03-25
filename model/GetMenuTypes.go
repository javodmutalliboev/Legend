package model

import "Legend/database"

func GetMenuTypes() ([]MenuType, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM menu_type ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuTypes []MenuType
	for rows.Next() {
		var menuType MenuType
		err := rows.Scan(&menuType.ID, &menuType.Title, &menuType.CreatedAt, &menuType.UpdatedAt)
		if err != nil {
			return nil, err
		}

		menuTypes = append(menuTypes, menuType)
	}

	return menuTypes, nil
}
