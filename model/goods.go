package model

import (
	"Legend/database"
	"fmt"
	"strings"

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

func GetGoods(menu_id, page, limit int) (*GoodsWrapper, error) {
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

	var goodsWrapper GoodsWrapper
	// in goods array, only get the goods that are in the current page
	for i := (page - 1) * limit; i < page*limit && i < len(goods); i++ {
		goodsWrapper.Goods = append(goodsWrapper.Goods, goods[i])
	}

	goodsWrapper.Count = len(goods)

	return &goodsWrapper, nil
}

// dive into menu.Children recursively. when children is null, get goods from db by the id of the current menu and append it to goods
func getGoods(menu *Menu, goods *[]Goods) error {
	if menu.Children == nil {
		// get goods from db by the id of the current menu and append it to goods
		db := database.DB()
		defer db.Close()

		rows, err := db.Query("SELECT id, menu_id, name, brand, sizes, price, discount, colors, description, created_at, updated_at FROM goods WHERE menu_id = $1 ORDER BY id", menu.ID)
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

	// extract Goods fields that are not empty
	var fields []string
	var args []interface{}
	i := 1

	if g.Name != "" {
		fields = append(fields, fmt.Sprintf("name = $%d", i))
		args = append(args, g.Name)
		i++
	}

	if g.Brand != "" {
		fields = append(fields, fmt.Sprintf("brand = $%d", i))
		args = append(args, g.Brand)
		i++
	}

	// ... repeat for other fields ...
	if len(g.Sizes) > 0 {
		// first remove empty strings from g.Sizes
		var sizes []string
		for _, size := range g.Sizes {
			if size != "" {
				sizes = append(sizes, size)
			}
		}
		g.Sizes = sizes

		fields = append(fields, fmt.Sprintf("sizes = $%d", i))
		args = append(args, pq.Array(g.Sizes))
		i++
	}

	if g.Price != 0 {
		fields = append(fields, fmt.Sprintf("price = $%d", i))
		args = append(args, g.Price)
		i++
	}

	fields = append(fields, fmt.Sprintf("discount = $%d", i))
	args = append(args, g.Discount)
	i++

	if len(g.Colors) > 0 {
		// first remove empty strings from g.Colors
		var colors []string
		for _, color := range g.Colors {
			if color != "" {
				colors = append(colors, color)
			}
		}
		g.Colors = colors

		fields = append(fields, fmt.Sprintf("colors = $%d", i))
		args = append(args, pq.Array(g.Colors))
		i++
	}

	if g.Description != "" {
		fields = append(fields, fmt.Sprintf("description = $%d", i))
		args = append(args, g.Description)
		i++
	}

	// build the SQL query
	sql := fmt.Sprintf("UPDATE goods SET %s, updated_at = NOW() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, g.ID)

	_, err := db.Exec(sql, args...)
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

func GetGoodsByID(id int64) (Goods, error) {
	db := database.DB()
	defer db.Close()

	var g Goods
	err := db.QueryRow("SELECT id, menu_id, name, brand, sizes, price, discount, colors, description, created_at, updated_at FROM goods WHERE id = $1", id).Scan(&g.ID, &g.MenuID, &g.Name, &g.Brand, pq.Array(&g.Sizes), &g.Price, &g.Discount, pq.Array(&g.Colors), &g.Description, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		return Goods{}, err
	}

	// get photos
	photos, err := GetGoodsPhotos(g.ID)
	if err != nil {
		return Goods{}, err
	}
	g.Photos = photos

	return g, nil
}

func GetHomeGoods(menu_type int) ([]Goods, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT g.id FROM goods g, menu m WHERE g.menu_id = m.id AND m.type = $1 ORDER BY RANDOM() LIMIT 4", menu_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goods []Goods
	for rows.Next() {
		var g Goods
		err = rows.Scan(&g.ID)
		if err != nil {
			return nil, err
		}

		// get photos
		photos, err := GetGoodsPhotos(g.ID)
		if err != nil {
			return nil, err
		}
		g.Photos = photos

		goods = append(goods, g)
	}

	return goods, nil
}

func GetRecommendedGoods(menu_type int) ([]Goods, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT g.id, g.name, g.price, g.discount FROM goods g, menu m WHERE g.menu_id = m.id AND m.type = $1 ORDER BY RANDOM() LIMIT 10", menu_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goods []Goods
	for rows.Next() {
		var g Goods
		err = rows.Scan(&g.ID, &g.Name, &g.Price, &g.Discount)
		if err != nil {
			return nil, err
		}

		// get photos
		photos, err := GetGoodsPhotos(g.ID)
		if err != nil {
			return nil, err
		}
		g.Photos = photos

		goods = append(goods, g)
	}

	return goods, nil
}

func SearchGoods(menu_type, page, limit int, keyword string) (*GoodsWrapper, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT g.id, g.menu_id, g.name, g.brand, g.sizes, g.price, g.discount, g.colors, g.description, g.created_at, g.updated_at FROM goods g, menu m WHERE g.menu_id = m.id AND m.type = $1 AND (g.name ILIKE $2 OR g.brand ILIKE $2 OR g.description ILIKE $2) ORDER BY g.id", menu_type, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goods []Goods
	for rows.Next() {
		var g Goods
		err = rows.Scan(&g.ID, &g.MenuID, &g.Name, &g.Brand, pq.Array(&g.Sizes), &g.Price, &g.Discount, pq.Array(&g.Colors), &g.Description, &g.CreatedAt, &g.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// get photos
		photos, err := GetGoodsPhotos(g.ID)
		if err != nil {
			return nil, err
		}
		g.Photos = photos

		goods = append(goods, g)
	}

	var goodsWrapper GoodsWrapper
	// in goods array, only get the goods that are in the current page
	for i := (page - 1) * limit; i < page*limit && i < len(goods); i++ {
		goodsWrapper.Goods = append(goodsWrapper.Goods, goods[i])
	}

	goodsWrapper.Count = len(goods)

	return &goodsWrapper, nil
}

func GetGoodsWithDiscount(menu_type, page, limit int) (*GoodsWrapper, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT g.id, g.menu_id, g.name, g.brand, g.sizes, g.price, g.discount, g.colors, g.description, g.created_at, g.updated_at FROM goods g, menu m WHERE g.menu_id = m.id AND m.type = $1 AND g.discount > 0 ORDER BY g.id", menu_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goods []Goods
	for rows.Next() {
		var g Goods
		err = rows.Scan(&g.ID, &g.MenuID, &g.Name, &g.Brand, pq.Array(&g.Sizes), &g.Price, &g.Discount, pq.Array(&g.Colors), &g.Description, &g.CreatedAt, &g.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// get photos
		photos, err := GetGoodsPhotos(g.ID)
		if err != nil {
			return nil, err
		}
		g.Photos = photos

		goods = append(goods, g)
	}

	var goodsWrapper GoodsWrapper
	// in goods array, only get the goods that are in the current page
	for i := (page - 1) * limit; i < page*limit && i < len(goods); i++ {
		goodsWrapper.Goods = append(goodsWrapper.Goods, goods[i])
	}

	goodsWrapper.Count = len(goods)

	return &goodsWrapper, nil
}

func GetMenuGoods(menu_id, page, limit int) (*GoodsWrapper, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT id FROM goods WHERE menu_id = $1 ORDER BY id", menu_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goods []Goods
	for rows.Next() {
		var g Goods
		err = rows.Scan(&g.ID)
		if err != nil {
			return nil, err
		}

		// get photos
		photos, err := GetGoodsPhotos(g.ID)
		if err != nil {
			return nil, err
		}
		g.Photos = photos

		goods = append(goods, g)
	}

	var goodsWrapper GoodsWrapper
	// in goods array, only get the goods that are in the current page
	for i := (page - 1) * limit; i < page*limit && i < len(goods); i++ {
		goodsWrapper.Goods = append(goodsWrapper.Goods, goods[i])
	}

	goodsWrapper.Count = len(goods)

	return &goodsWrapper, nil
}
