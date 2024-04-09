package model

import (
	"Legend/database"
	"errors"
	"fmt"
	"strings"
)

type CTWInformation struct {
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

func CreateCTWInformation(ctw *CTWInformation) error {
	db := database.DB()
	defer db.Close()

	// check for at least one ctw_information row exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM ctw_information)").Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// return error of 'ctw information already exists'
		return errors.New("ctw information already exists")
	}

	_, err = db.Exec("INSERT INTO ctw_information (heading_uz, heading_ru, heading_en, description_uz, description_ru, description_en) VALUES ($1, $2, $3, $4, $5, %6)", ctw.HeadingUz, ctw.HeadingRu, ctw.HeadingEn, ctw.DescriptionUz, ctw.DescriptionRu, ctw.DescriptionEn)
	if err != nil {
		return err
	}

	return nil
}

func GetCTWInformation() (CTWInformation, error) {
	db := database.DB()
	defer db.Close()

	var ctw CTWInformation
	err := db.QueryRow("SELECT * FROM ctw_information").Scan(&ctw.ID, &ctw.HeadingUz, &ctw.HeadingRu, &ctw.HeadingEn, &ctw.DescriptionUz, &ctw.DescriptionRu, &ctw.DescriptionEn, &ctw.CreatedAt, &ctw.UpdatedAt)
	if err != nil {
		return ctw, err
	}

	return ctw, nil
}

func UpdateCTWInformation(ctw *CTWInformation) error {
	db := database.DB()
	defer db.Close()

	var fields []string
	var args []interface{}
	i := 1

	if ctw.HeadingUz != "" {
		fields = append(fields, fmt.Sprintf("heading_uz = $%d", i))
		args = append(args, ctw.HeadingUz)
		i++
	}

	if ctw.HeadingRu != "" {
		fields = append(fields, fmt.Sprintf("heading_ru = $%d", i))
		args = append(args, ctw.HeadingRu)
		i++
	}

	if ctw.HeadingEn != "" {
		fields = append(fields, fmt.Sprintf("heading_en = $%d", i))
		args = append(args, ctw.HeadingEn)
		i++
	}

	if ctw.DescriptionUz != "" {
		fields = append(fields, fmt.Sprintf("description_uz = $%d", i))
		args = append(args, ctw.DescriptionUz)
		i++
	}

	if ctw.DescriptionRu != "" {
		fields = append(fields, fmt.Sprintf("description_ru = $%d", i))
		args = append(args, ctw.DescriptionRu)
		i++
	}

	if ctw.DescriptionEn != "" {
		fields = append(fields, fmt.Sprintf("description_en = $%d", i))
		args = append(args, ctw.DescriptionEn)
		i++
	}

	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	// build the SQL query
	sql := fmt.Sprintf("UPDATE ctw_information SET %s, updated_at = NOW() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, ctw.ID)

	_, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCTWInformation(id int) error {
	db := database.DB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM ctw_information WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
