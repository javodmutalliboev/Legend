package model

import "Legend/database"

type Order struct {
	ID               int64        `json:"id"`
	Goods            []OrderGoods `json:"goods"`
	CustomerName     string       `json:"customer_name"`
	CustomerSurname  string       `json:"customer_surname"`
	CustomerRegion   string       `json:"customer_region"`
	CustomerDistrict string       `json:"customer_district"`
	CustomerAddress  string       `json:"customer_address"`
	CustomerPhone    string       `json:"customer_phone"`
	CustomerPhone2   string       `json:"customer_phone2"`
	Canceled         bool         `json:"canceled"`
	Delivered        bool         `json:"delivered"`
	CreatedAt        string       `json:"created_at"`
}

func (o *Order) Create() (*int64, error) {
	db := database.DB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(`INSERT INTO "order" (customer_name, customer_surname, customer_region, customer_district, customer_address, customer_phone, customer_phone2) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, o.CustomerName, o.CustomerSurname, o.CustomerRegion, o.CustomerDistrict, o.CustomerAddress, o.CustomerPhone, o.CustomerPhone2).Scan(&o.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, og := range o.Goods {
		_, err = tx.Exec("INSERT INTO order_goods (goods_id, order_id, color, size, quantity) VALUES ($1, $2, $3, $4, $5)", og.GoodsID, o.ID, og.Color, og.Size, og.Quantity)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &o.ID, nil
}

func GetOrders(delivered bool) ([]*Order, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query(`SELECT id, customer_name, customer_surname, customer_region, customer_district, customer_address, customer_phone, customer_phone2, canceled, delivered, created_at FROM "order" WHERE delivered = $1 ORDER BY created_at DESC`, delivered)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []*Order{}
	for rows.Next() {
		var o Order
		err = rows.Scan(&o.ID, &o.CustomerName, &o.CustomerSurname, &o.CustomerRegion, &o.CustomerDistrict, &o.CustomerAddress, &o.CustomerPhone, &o.CustomerPhone2, &o.Canceled, &o.Delivered, &o.CreatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &o)
	}

	return orders, nil
}

func (o *Order) GetGoods() error {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query(`SELECT og.goods_id, og.order_id, og.color, og.size, og.quantity, og.created_at, g.name_uz, g.name_ru, g.name_en, g.brand_uz, g.brand_ru, g.brand_en, g.price, g.discount FROM order_goods og JOIN goods g ON og.goods_id = g.id WHERE og.order_id = $1`, o.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	o.Goods = []OrderGoods{}
	for rows.Next() {
		var og OrderGoods
		err = rows.Scan(&og.GoodsID, &og.OrderID, &og.Color, &og.Size, &og.Quantity, &og.CreatedAt, &og.NameUz, &og.NameRu, &og.NameEn, &og.BrandUz, &og.BrandRu, &og.BrandEn, &og.Price, &og.Discount)
		if err != nil {
			return err
		}

		photos, err := GetGoodsPhotos(og.GoodsID)
		if err != nil {
			return err
		}
		og.Photos = photos

		o.Goods = append(o.Goods, og)
	}

	return nil
}

func GetOrder(id int64) (*Order, error) {
	db := database.DB()
	defer db.Close()

	o := &Order{}
	err := db.QueryRow(`SELECT id, customer_name, customer_surname, customer_region, customer_district, customer_address, customer_phone, customer_phone2, canceled, delivered, created_at FROM "order" WHERE id = $1`, id).Scan(&o.ID, &o.CustomerName, &o.CustomerSurname, &o.CustomerRegion, &o.CustomerDistrict, &o.CustomerAddress, &o.CustomerPhone, &o.CustomerPhone2, &o.Canceled, &o.Delivered, &o.CreatedAt)
	if err != nil {
		return nil, err
	}

	err = o.GetGoods()
	if err != nil {
		return nil, err
	}

	return o, nil
}
