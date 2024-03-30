package model

type GoodsWrapper struct {
	Goods []Goods `json:"goods"`
	Count int     `json:"count"`
}
