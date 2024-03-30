package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func GetGoodsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		goods_id, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid goods_id").Send(w)
			return
		}

		// get the goods
		goods, err := model.GetGoodsByID(int64(goods_id))
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, goods).Send(w)
	}
}
