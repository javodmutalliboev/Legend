package model

type Menu struct {
	ID        int    `json:"id"`
	ParentID  *int   `json:"parent_id"`
	TitleUz   string `json:"title_uz"`
	TitleRu   string `json:"title_ru"`
	TitleEn   string `json:"title_en"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Type      int    `json:"type"`
	Children  []Menu `json:"children"`
}
