package model

type OrderGoods struct {
	Photos    []GoodsPhoto `json:"photos"`
	NameUz    string       `json:"name_uz"`
	NameRu    string       `json:"name_ru"`
	NameEn    string       `json:"name_en"`
	BrandUz   string       `json:"brand_uz"`
	BrandRu   string       `json:"brand_ru"`
	BrandEn   string       `json:"brand_en"`
	Price     float64      `json:"price"`
	Discount  float64      `json:"discount"`
	GoodsID   int64        `json:"goods_id"`
	OrderID   int64        `json:"order_id"`
	Color     string       `json:"color"`
	Size      string       `json:"size"`
	Quantity  int          `json:"quantity"`
	CreatedAt string       `json:"created_at"`
}
