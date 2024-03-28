package admin

import (
	"Legend/model"
	"Legend/response"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

func CreateGoods() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request body is multipart/form-data
		//  max file size is 100MB
		// no max number of files
		// now parse form without comment
		err := r.ParseMultipartForm(100 << 20) // 100MB
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		// get the form data
		formData := r.MultipartForm

		// declare a new goods
		var goods model.Goods

		// get the form data
		menu_id_str := formData.Value["menu_id"][0]
		menu_id, err := strconv.Atoi(menu_id_str)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid menu_id").Send(w)
			return
		}
		goods.MenuID = menu_id

		goods.Name = formData.Value["name"][0]

		goods.Brand = formData.Value["brand"][0]

		// photos
		photos := formData.File["photo"]
		for _, photo := range photos {
			// check if the photo is of type image. It can be any type of image
			if !isImage(photo) {
				log.Printf("%s: %s", r.URL.Path, "Invalid image file")
				response.NewResponse("error", http.StatusBadRequest, "Invalid image file").Send(w)
				return
			}

			// check if the photo is not more than 16MB
			if photo.Size > 16<<20 { // 16MB
				log.Printf("%s: %s", r.URL.Path, "Image file too large")
				response.NewResponse("error", http.StatusBadRequest, "Image file too large").Send(w)
				return
			}
		}

		goods.Sizes = formData.Value["size"]

		priceStr := formData.Value["price"][0]
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid price").Send(w)
			return
		}
		goods.Price = price

		discountStr := formData.Value["discount"][0]
		discount, err := strconv.ParseFloat(discountStr, 64)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid discount").Send(w)
			return
		}
		goods.Discount = discount

		goods.Colors = formData.Value["color"]

		goods.Description = formData.Value["description"][0]

		// create the goods
		goods_id, err := model.CreateGoods(&goods)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		// save the photos
		for _, photo := range photos {
			// save the photo
			err := model.SavePhoto(photo, goods_id)
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err)

				// delete goods
				err = model.DeleteGoods(goods_id)
				if err != nil {
					log.Printf("%s: %s", r.URL.Path, err)
				}

				response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
				return
			}
		}

		response.NewResponse("success", http.StatusCreated, "Goods created").Send(w)
	}
}

func isImage(file *multipart.FileHeader) bool {
	// open the file
	f, err := file.Open()
	if err != nil {
		return false
	}
	defer f.Close()

	// get the first 512 bytes
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return false
	}

	// check if the file is of type image
	contentType := http.DetectContentType(buffer)

	return strings.HasPrefix(contentType, "image/")
}