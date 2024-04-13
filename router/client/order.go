package client

import (
	"Legend/interface_package"
	"Legend/model"
	"Legend/response"
	"encoding/json"
	"log"
	"net/http"
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

		err = order.Create()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err.Error())
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "Order created").Send(w)
	}
}
