package model

import "Legend/database"

func GetMenuTypes() ([]MenuType, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT id, title_uz, title_ru, title_en, created_at, updated_at FROM menu_type ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuTypes []MenuType
	for rows.Next() {
		var menuType MenuType
		err := rows.Scan(&menuType.ID, &menuType.TitleUz, &menuType.TitleRu, &menuType.TitleEn, &menuType.CreatedAt, &menuType.UpdatedAt)
		if err != nil {
			return nil, err
		}

		menuTypes = append(menuTypes, menuType)
	}

	return menuTypes, nil
}
