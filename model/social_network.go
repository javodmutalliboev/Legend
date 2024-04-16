package model

import "Legend/database"

type SocialNetwork struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      []byte `json:"icon"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (sn *SocialNetwork) Create() error {
	db := database.DB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO social_network (name, icon, url) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sn.Name, sn.Icon, sn.URL)
	if err != nil {
		return err
	}

	return nil
}
