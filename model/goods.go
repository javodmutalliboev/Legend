package model

import (
	"Legend/database"

	"github.com/lib/pq"
)

type Goods struct {
	ID          int64        `json:"id"`
	MenuID      int          `json:"menu_id"`
	Name        string       `json:"name"`
	Brand       string       `json:"brand"`
	Photos      []GoodsPhoto `json:"photos"`
	Sizes       []string     `json:"sizes"`
	Price       float64      `json:"price"`
	Discount    float64      `json:"discount"`
	Colors      []string     `json:"colors"`
	Description string       `json:"description"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

func CreateGoods(g *Goods) (int64, error) {
	db := database.DB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO goods (menu_id, name, brand, sizes, price, discount, colors, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var goods_id int64
	err = stmt.QueryRow(g.MenuID, g.Name, g.Brand, pq.Array(g.Sizes), g.Price, g.Discount, pq.Array(g.Colors), g.Description).Scan(&goods_id)
	if err != nil {
		return 0, err
	}

	return goods_id, nil
}

func DeleteGoods(id int64) error {
	db := database.DB()
	defer db.Close()

	// delete photos
	err := deleteGoodsPhotos(id)
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM goods WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func GetGoods(menu_id int) ([]Goods, error) {
	db := database.DB()
	defer db.Close()

	menu, err := GetMenu(menu_id)
	if err != nil {
		return nil, err
	}

	var goods []Goods

	err = getGoods(menu, &goods)
	if err != nil {
		return nil, err
	}

	return goods, nil
}

// dive into menu.Children recursively. when children is null, get goods from db by the id of the current menu and append it to goods
func getGoods(menu *Menu, goods *[]Goods) error {
	if menu.Children == nil {
		// get goods from db by the id of the current menu and append it to goods
		db := database.DB()
		defer db.Close()

		rows, err := db.Query("SELECT id, menu_id, name, brand, sizes, price, discount, colors, description, created_at, updated_at FROM goods WHERE menu_id = $1", menu.ID)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var g Goods
			err = rows.Scan(&g.ID, &g.MenuID, &g.Name, &g.Brand, pq.Array(&g.Sizes), &g.Price, &g.Discount, pq.Array(&g.Colors), &g.Description, &g.CreatedAt, &g.UpdatedAt)
			if err != nil {
				return err
			}

			// get photos
			photos, err := GetGoodsPhotos(g.ID)
			if err != nil {
				return err
			}
			g.Photos = photos

			*goods = append(*goods, g)
		}

		return nil
	}

	for _, child := range menu.Children {
		getGoods(&child, goods)
	}

	return nil
}

func UpdateGoods(g *Goods) error {
	db := database.DB()
	defer db.Close()

	_, err := db.Exec("UPDATE goods SET name = $1, brand = $2, sizes = $3, price = $4, discount = $5, colors = $6, description = $7, updated_at = NOW() WHERE id = $8", g.Name, g.Brand, pq.Array(g.Sizes), g.Price, g.Discount, pq.Array(g.Colors), g.Description, g.ID)
	if err != nil {
		return err
	}

	return nil
}

func deleteGoodsPhotos(goods_id int64) error {
	db := database.DB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM goods_photo WHERE goods_id = $1", goods_id)
	if err != nil {
		return err
	}

	return nil
}

func GetPhoto(id int64) ([]byte, error) {
	db := database.DB()
	defer db.Close()

	var photo []byte
	err := db.QueryRow("SELECT photo FROM goods_photo WHERE id = $1", id).Scan(&photo)
	if err != nil {
		return nil, err
	}

	return photo, nil
}
