package model

type Admin struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
