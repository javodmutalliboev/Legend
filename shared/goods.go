package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetGoods() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the menu_id
		menu_id_str := mux.Vars(r)["menu_id"]
		menu_id, err := strconv.Atoi(menu_id_str)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid menu_id").Send(w)
			return
		}

		// get all goods of a menu
		goods, err := model.GetGoods(menu_id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, goods).Send(w)
	}
}

func GetGoodsPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		photo_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid photo_id").Send(w)
			return
		}

		// get the photo
		photo, err := model.GetPhoto(int64(photo_id))
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		contentType := http.DetectContentType(photo)

		// send the photo
		w.Header().Set("Content-Type", contentType)
		w.Write(photo)
	}
}
