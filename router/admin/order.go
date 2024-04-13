package admin

import (
	"Legend/interface_package"
	"Legend/model"
	"Legend/response"
	"Legend/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

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

		delivered_str := r.URL.Query().Get("delivered")
		delivered, err := strconv.ParseBool(delivered_str)
		if err != nil {
			delivered = false
		}

		var ordersPtr []*model.Order
		ordersPtr, err = model.GetOrders(delivered)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		data := map[string]interface{}{
			"count": len(ordersPtr),
		}

		ordersPtr = utils.SliceByPageLimit(ordersPtr, page, limit)

		for _, orderPtr := range ordersPtr {
			err = orderPtr.GetGoods()
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err.Error())
				response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
				return
			}
		}

		orders := make([]model.Order, len(ordersPtr))
		for i, orderPtr := range ordersPtr {
			orders[i] = *orderPtr
		}

		data["orders"] = orders

		response.NewResponse("success", http.StatusOK, data).Send(w)
	}
}

func UpdateOrderCanceled() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order interface_package.Order = &model.Order{}
		err := json.NewDecoder(r.Body).Decode(order)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusBadRequest, "Invalid request body").Send(w)
			return
		}

		err = order.ToggleCanceled()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Order canceled status updated").Send(w)
	}
}

func UpdateOrderDelivered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order interface_package.Order = &model.Order{}
		err := json.NewDecoder(r.Body).Decode(order)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusBadRequest, "Invalid request body").Send(w)
			return
		}

		err = order.ToggleDelivered()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Order delivered status updated").Send(w)
	}
}
