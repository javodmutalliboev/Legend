package model

type OrderGoods struct {
	GoodsID   int64  `json:"goods_id"`
	OrderID   int64  `json:"order_id"`
	Color     string `json:"color"`
	Size      string `json:"size"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}
