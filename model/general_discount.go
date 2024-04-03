package model

import "Legend/database"

type GeneralDiscount struct {
	ID        int     `json:"id"`
	MenuType  int     `json:"menu_type"`
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func CreateGeneralDiscount(gd *GeneralDiscount) error {
	db := database.DB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO general_discount (menu_type, value, unit) VALUES ($1, $2, $3)", gd.MenuType, gd.Value, gd.Unit)
	if err != nil {
		return err
	}

	return nil
}

func GetGeneralDiscountByMenuType(menuType int) (bool, error) {
	db := database.DB()
	defer db.Close()

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM general_discount WHERE menu_type = $1)", menuType).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
