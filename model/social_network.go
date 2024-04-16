package model

import (
	"Legend/database"
	"fmt"
	"strings"
)

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

func GetSocialNetworks() ([]SocialNetwork, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, url, created_at, updated_at FROM social_network")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	socialNetworks := []SocialNetwork{}
	for rows.Next() {
		sn := SocialNetwork{}
		err := rows.Scan(&sn.ID, &sn.Name, &sn.URL, &sn.CreatedAt, &sn.UpdatedAt)
		if err != nil {
			return nil, err
		}
		socialNetworks = append(socialNetworks, sn)
	}

	return socialNetworks, nil
}

func GetSocialNetwork(id int) (SocialNetwork, error) {
	db := database.DB()
	defer db.Close()

	sn := SocialNetwork{}
	err := db.QueryRow("SELECT id, name, url, created_at, updated_at FROM social_network WHERE id = $1", id).Scan(&sn.ID, &sn.Name, &sn.URL, &sn.CreatedAt, &sn.UpdatedAt)
	if err != nil {
		return SocialNetwork{}, err
	}

	return sn, nil
}

func GetSocialNetworkIcon(id int) (*[]byte, error) {
	db := database.DB()
	defer db.Close()

	var icon []byte
	err := db.QueryRow("SELECT icon FROM social_network WHERE id = $1", id).Scan(&icon)
	if err != nil {
		return nil, err
	}

	return &icon, nil
}

func (sn *SocialNetwork) Update() error {
	db := database.DB()
	defer db.Close()

	var fields []string
	var args []interface{}
	i := 1

	if sn.Name != "" {
		fields = append(fields, fmt.Sprintf("name = $%d", i))
		args = append(args, sn.Name)
		i++
	}

	if sn.Icon != nil {
		fields = append(fields, fmt.Sprintf("icon = $%d", i))
		args = append(args, sn.Icon)
		i++
	}

	if sn.URL != "" {
		fields = append(fields, fmt.Sprintf("url = $%d", i))
		args = append(args, sn.URL)
		i++
	}

	if len(fields) == 0 {
		return fmt.Errorf("no values to update")
	}

	// build the SQL query
	query := fmt.Sprintf("UPDATE social_network SET %s, updated_at = now() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, sn.ID)

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSocialNetwork(id int) error {
	db := database.DB()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM social_network WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
