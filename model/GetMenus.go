package model

import "Legend/database"

func GetMenus() ([]Menu, error) {
	database := database.DB()
	defer database.Close()

	// get menus with its children taken place in their corresponding parent
	rows, err := database.Query("SELECT * FROM menu WHERE parent_id IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []Menu
	for rows.Next() {
		var menu Menu
		err := rows.Scan(&menu.ID, &menu.ParentID, &menu.Title, &menu.CreatedAt, &menu.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// get children
		err = getChildren(&menu)
		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func getChildren(m *Menu) error {
	database := database.DB()
	defer database.Close()

	rows, err := database.Query("SELECT * FROM menu WHERE parent_id = $1", m.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var menu Menu
		err := rows.Scan(&menu.ID, &menu.ParentID, &menu.Title, &menu.CreatedAt, &menu.UpdatedAt)
		if err != nil {
			return err
		}

		// get children
		err = getChildren(&menu)
		if err != nil {
			return err
		}

		m.Children = append(m.Children, menu)
	}

	return nil
}
