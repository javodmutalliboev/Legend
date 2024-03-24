package model

import "Legend/database"

type Menu struct {
	ID        int    `json:"id"`
	ParentID  *int   `json:"parent_id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Children  []Menu `json:"children"`
}

func (m *Menu) MenuSave() error {
	// Save the menu to the database
	// Open a database connection
	// Insert the menu to the database
	// Close the database connection
	database := database.DB()
	defer database.Close()

	// save title to the database
	_, err := database.Exec("INSERT INTO menu (title) VALUES ($1)", m.Title)
	if err != nil {
		return err
	}

	return nil
}

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
