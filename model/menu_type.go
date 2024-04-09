package model

type MenuType struct {
	ID        int    `json:"id"`
	TitleUz   string `json:"title_uz"`
	TitleRu   string `json:"title_ru"`
	TitleEn   string `json:"title_en"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
