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

	// GET /general_discount/{menu_type:[0-9]+} will return a general discount
	router.HandleFunc("/general_discount/{menu_type:[0-9]+}", shared.GetGeneralDiscount()).Methods("GET")

	// GET /legend_information will return legend_information
	router.HandleFunc("/legend_information", shared.GetLegendInformation()).Methods("GET")

	// GET /ctw_information will return ctw_information
	router.HandleFunc("/ctw_information", shared.GetCTWInformation()).Methods("GET")

	// GET /home/goods/{menu_type:[0-9]+} will return home goods
	router.HandleFunc("/home/goods/{menu_type:[0-9]+}", shared.GetHomeGoods()).Methods("GET")

	return router
}
