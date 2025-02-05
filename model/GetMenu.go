package model

import "Legend/database"

func GetMenu(id int) (*Menu, error) {
	db := database.DB()
	defer db.Close()

	var menu Menu
	err := db.QueryRow("SELECT id, parent_id, title_uz, title_ru, title_en, created_at, updated_at, type FROM menu WHERE id = $1", id).Scan(&menu.ID, &menu.ParentID, &menu.TitleUz, &menu.TitleRu, &menu.TitleEn, &menu.CreatedAt, &menu.UpdatedAt, &menu.Type)
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
