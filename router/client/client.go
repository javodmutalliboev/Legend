package client

import (
	"Legend/shared"

	"github.com/gorilla/mux"
)

func ClientRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1/client").Subrouter()

	// GET /goods/{menu_id:[0-9]+} will return all goods of a menu
	router.HandleFunc("/goods/{menu_id:[0-9]+}", shared.GetGoods()).Methods("GET")

	return router
}
