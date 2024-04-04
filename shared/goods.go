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
		// page and limit
		page_str := r.URL.Query().Get("page")
		limit_str := r.URL.Query().Get("limit")
		page, err := strconv.Atoi(page_str)
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(limit_str)
		if err != nil {
			limit = 10
		}

		// get the menu_id
		menu_id_str := mux.Vars(r)["menu_id"]
		menu_id, err := strconv.Atoi(menu_id_str)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid menu_id").Send(w)
			return
		}

		// get all goods of a menu
		goodsWrapper, err := model.GetGoods(menu_id, page, limit)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, *goodsWrapper).Send(w)
	}
}

func GetHomeGoods() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the menu_id
		menu_type_str := mux.Vars(r)["menu_type"]
		menu_type, err := strconv.Atoi(menu_type_str)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid menu_id").Send(w)
			return
		}

		// get home goods
		goods, err := model.GetHomeGoods(menu_type)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, goods).Send(w)
	}
}
