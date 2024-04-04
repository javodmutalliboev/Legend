package model

import (
	"Legend/database"
	"errors"
	"fmt"
	"strings"
)

type GeneralDiscount struct {
	ID        int     `json:"id"`
	MenuType  int     `json:"menu_type"`
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Title     string  `json:"title"`
}

func CreateGeneralDiscount(gd *GeneralDiscount) error {
	db := database.DB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO general_discount (menu_type, title, value, unit) VALUES ($1, $2, $3, $4)", gd.MenuType, gd.Title, gd.Value, gd.Unit)
	if err != nil {
		return err
	}

	return nil
}

func CheckGeneralDiscountExistenceByMenuType(menuType int) (bool, error) {
	db := database.DB()
	defer db.Close()

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM general_discount WHERE menu_type = $1)", menuType).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func GetGeneralDiscountByMenuType(menuType int) (GeneralDiscount, error) {
	db := database.DB()
	defer db.Close()

	var gd GeneralDiscount
	err := db.QueryRow("SELECT * FROM general_discount WHERE menu_type = $1", menuType).Scan(&gd.ID, &gd.MenuType, &gd.Value, &gd.Unit, &gd.CreatedAt, &gd.UpdatedAt, &gd.Title)
	if err != nil {
		return GeneralDiscount{}, err
	}

	return gd, nil
}

func UpdateGeneralDiscount(gd *GeneralDiscount) error {
	db := database.DB()
	defer db.Close()

	var fields []string
	var args []interface{}
	i := 1

	if gd.Title != "" {
		fields = append(fields, fmt.Sprintf("title = $%d", i))
		args = append(args, gd.Title)
		i++
	}

	fields = append(fields, fmt.Sprintf("value = $%d", i))
	args = append(args, gd.Value)
	i++

	if gd.Unit != "" {
		fields = append(fields, fmt.Sprintf("unit = $%d", i))
		args = append(args, gd.Unit)
		i++
	}

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	// build the SQL query
	sql := fmt.Sprintf("UPDATE general_discount SET %s, updated_at = NOW() WHERE menu_type = $%d", strings.Join(fields, ", "), i)
	args = append(args, gd.MenuType)

	_, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteGeneralDiscount(menuType int) error {
	db := database.DB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM general_discount WHERE menu_type = $1", menuType)
	if err != nil {
		return err
	}

	return nil
}
