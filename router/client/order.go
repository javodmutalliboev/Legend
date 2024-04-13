package client

import (
	"Legend/interface_package"
	"Legend/model"
	"Legend/response"
	"Legend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order interface_package.Order = &model.Order{}

		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		id, err := order.Create()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, map[string]any{
			"id":      *id,
			"message": "Order created",
		}).Send(w)
	}
}

func GetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page_str := r.URL.Query().Get("page")
		page, err := strconv.Atoi(page_str)
		if err != nil {
			page = 1
		}

		limit_str := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(limit_str)
		if err != nil {
			limit = 10
		}

		ids_str, ok := r.URL.Query()["id"]
		if !ok || len(ids_str) < 1 {
			log.Printf("%s: %s", r.URL.Path, "No 'id' parameter in the query string")
			response.NewResponse("error", http.StatusBadRequest, "No 'id' parameter in the query string").Send(w)
			return
		}

		ids := make([]int64, len(ids_str))
		for i, id_str := range ids_str {
			id, err := strconv.ParseInt(id_str, 10, 64)
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err.Error())
				response.NewResponse("error", http.StatusBadRequest, fmt.Sprintf("Invalid 'id' parameter: '%s'", id_str)).Send(w)
				return
			}
			ids[i] = id
		}

		var orders []model.Order
		for _, id := range ids {
			order, err := model.GetOrder(id)
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err.Error())
				response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
				return
			}
			orders = append(orders, *order)
		}

		data := map[string]any{
			"count": len(orders),
		}

		orders = utils.SliceByPageLimit(orders, page, limit)

		data["orders"] = orders

		response.NewResponse("success", http.StatusOK, data).Send(w)
	}
}
