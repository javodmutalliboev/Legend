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

func (o *Order) Create() error {
	db := database.DB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow(`INSERT INTO "order" (customer_name, customer_surname, customer_region, customer_district, customer_address, customer_phone, customer_phone2) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`, o.CustomerName, o.CustomerSurname, o.CustomerRegion, o.CustomerDistrict, o.CustomerAddress, o.CustomerPhone, o.CustomerPhone2).Scan(&o.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, og := range o.Goods {
		_, err = tx.Exec("INSERT INTO order_goods (goods_id, order_id, color, size, quantity) VALUES ($1, $2, $3, $4, $5)", og.GoodsID, o.ID, og.Color, og.Size, og.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
