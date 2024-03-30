package client

import (
	"Legend/shared"

	"github.com/gorilla/mux"
)

func ClientRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1/client").Subrouter()

	// GET /menus/{type:[0-9]+} will return all menus of a type
	router.HandleFunc("/menus/{type:[0-9]+}", shared.GetMenus()).Methods("GET")

	// GET /goods/{menu_id:[0-9]+} will return all goods of a menu
	router.HandleFunc("/goods/{menu_id:[0-9]+}", shared.GetGoods()).Methods("GET")

	// GET /goods/photo/{id:[0-9]+} will return a photo of a goods
	router.HandleFunc("/goods/photo/{id:[0-9]+}", shared.GetGoodsPhoto()).Methods("GET")

	// GET /goods/{id:[0-9]+} will return a goods
	router.HandleFunc("/goods/id/{id:[0-9]+}", shared.GetGoodsByID()).Methods("GET")

	return router
}
