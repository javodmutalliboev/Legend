package model

type GoodsPhoto struct {
	ID        int64  `json:"id"`
	GoodsID   int64  `json:"goods_id"`
	Photo     []byte `json:"photo"`
	CreatedAt string `json:"created_at"`
}
