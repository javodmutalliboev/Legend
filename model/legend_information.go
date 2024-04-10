package model

import (
	"Legend/database"
	"errors"
	"fmt"
	"strings"
)

type LegendInformation struct {
	ID            int    `json:"id"`
	HeadingUz     string `json:"heading_uz"`
	HeadingRu     string `json:"heading_ru"`
	HeadingEn     string `json:"heading_en"`
	DescriptionUz string `json:"description_uz"`
	DescriptionRu string `json:"description_ru"`
	DescriptionEn string `json:"description_en"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func CreateLegendInformation(li *LegendInformation) error {
	db := database.DB()
	defer db.Close()

	// check for at least one legend_information row exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM legend_information)").Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// return error of 'legend information already exists'
		return errors.New("legend information already exists")
	}

	_, err = db.Exec("INSERT INTO legend_information (heading_uz, heading_ru, heading_en, description_uz, description_ru, description_en) VALUES ($1, $2, $3, $4, $5, $6)", li.HeadingUz, li.HeadingRu, li.HeadingEn, li.DescriptionUz, li.DescriptionRu, li.DescriptionEn)
	if err != nil {
		return err
	}

	return nil
}

func GetLegendInformation() (LegendInformation, error) {
	db := database.DB()
	defer db.Close()

	var li LegendInformation
	err := db.QueryRow("SELECT id, heading_uz, heading_ru, heading_en, description_uz, description_ru, description_en, created_at, updated_at FROM legend_information").Scan(&li.ID, &li.HeadingUz, &li.HeadingRu, &li.HeadingEn, &li.DescriptionUz, &li.DescriptionRu, &li.DescriptionEn, &li.CreatedAt, &li.UpdatedAt)
	if err != nil {
		return li, err
	}

	return li, nil
}

func UpdateLegendInformation(li *LegendInformation) error {
	db := database.DB()
	defer db.Close()

	var fields []string
	var args []interface{}
	i := 1

	if li.HeadingUz != "" {
		fields = append(fields, fmt.Sprintf("heading_uz = $%d", i))
		args = append(args, li.HeadingUz)
		i++
	}

	if li.HeadingRu != "" {
		fields = append(fields, fmt.Sprintf("heading_ru = $%d", i))
		args = append(args, li.HeadingRu)
		i++
	}

	if li.HeadingEn != "" {
		fields = append(fields, fmt.Sprintf("heading_en = $%d", i))
		args = append(args, li.HeadingEn)
		i++
	}

	if li.DescriptionUz != "" {
		fields = append(fields, fmt.Sprintf("description_uz = $%d", i))
		args = append(args, li.DescriptionUz)
		i++
	}

	if li.DescriptionRu != "" {
		fields = append(fields, fmt.Sprintf("description_ru = $%d", i))
		args = append(args, li.DescriptionRu)
		i++
	}

	if li.DescriptionEn != "" {
		fields = append(fields, fmt.Sprintf("description_en = $%d", i))
		args = append(args, li.DescriptionEn)
		i++
	}

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	// build the SQL query
	sql := fmt.Sprintf("UPDATE legend_information SET %s, updated_at = NOW() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, li.ID)

	_, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteLegendInformation(id int) error {
	db := database.DB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM legend_information WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
