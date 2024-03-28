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

func GetGoodsPhotos(goods_id int64) ([]GoodsPhoto, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT id, goods_id, created_at FROM goods_photo WHERE goods_id = $1", goods_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []GoodsPhoto
	for rows.Next() {
		var photo GoodsPhoto
		err := rows.Scan(&photo.ID, &photo.GoodsID, &photo.CreatedAt)
		if err != nil {
			return nil, err
		}

		photos = append(photos, photo)
	}

	return photos, nil
}

func DeletePhoto(id int64) error {
	db := database.DB()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM goods_photo WHERE id = $1")
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
