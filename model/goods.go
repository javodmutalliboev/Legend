package model

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
