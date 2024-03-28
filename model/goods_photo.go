package model

import (
	"Legend/database"
	"io"
	"mime/multipart"
)

type GoodsPhoto struct {
	ID        int64  `json:"id"`
	GoodsID   int64  `json:"goods_id"`
	Photo     []byte `json:"photo"`
	CreatedAt string `json:"created_at"`
}

func SavePhoto(photo *multipart.FileHeader, goods_id int64) error {
	// open the file
	f, err := photo.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	// get the file content
	bytea, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// save the photo
	db := database.DB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO goods_photo (goods_id, photo) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(goods_id, bytea)
	if err != nil {
		return err
	}

	return nil
}
