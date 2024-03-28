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

	_, err := db.Exec("DELETE FROM goods WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
