package model

type Menu struct {
	ID        int    `json:"id"`
	ParentID  *int   `json:"parent_id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Children  []Menu `json:"children"`
}
