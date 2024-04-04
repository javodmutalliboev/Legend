package admin

import (
	"Legend/middleware"
	"Legend/response"
	"Legend/shared"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AdminRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1/admin").Subrouter()

	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: someone is requesting me", r.URL.Path)
		json.NewEncoder(w).Encode(response.Response{Status: "success", Code: http.StatusOK, Data: "admin route"})
	}).Methods("GET")

	router.HandleFunc("/login", Login()).Methods("POST")
	router.HandleFunc("/logout", middleware.Chain(Logout(), middleware.Auth())).Methods("GET")

	// POST /menu will create a new menu
	router.HandleFunc("/menu", middleware.Chain(CreateMenu(), middleware.Auth())).Methods("POST")

	// POST /menu/sub will create a new sub menu
	router.HandleFunc("/menu/sub", middleware.Chain(CreateSubMenu(), middleware.Auth())).Methods("POST")

	// GET /menus/{type:[0-9]+} will return all menus of a type
	router.HandleFunc("/menus/{type:[0-9]+}", middleware.Chain(shared.GetMenus(), middleware.Auth())).Methods("GET")

	// PATCH /menu will update a menu
	router.HandleFunc("/menu", middleware.Chain(UpdateMenu(), middleware.Auth())).Methods("PATCH")

	// DELETE /menu/{id:[0-9]+} will delete a menu
	router.HandleFunc("/menu/{id:[0-9]+}", middleware.Chain(DeleteMenu(), middleware.Auth())).Methods("DELETE")

	// GET /menu/types will return all menu types
	router.HandleFunc("/menu/types", middleware.Chain(GetMenuTypes(), middleware.Auth())).Methods("GET")

	// POST /goods will create a new goods
	router.HandleFunc("/goods", middleware.Chain(CreateGoods(), middleware.Auth())).Methods("POST")

	// GET /goods/{menu_id:[0-9]+} will return all goods of a menu
	router.HandleFunc("/goods/{menu_id:[0-9]+}", middleware.Chain(shared.GetGoods(), middleware.Auth())).Methods("GET")

	// PATCH /goods will update a goods
	router.HandleFunc("/goods", middleware.Chain(UpdateGoods(), middleware.Auth())).Methods("PATCH")

	// DELETE /goods/{id:[0-9]+} will delete a goods
	router.HandleFunc("/goods/{id:[0-9]+}", middleware.Chain(DeleteGoods(), middleware.Auth())).Methods("DELETE")

	// GET /goods/photo/{id:[0-9]+} will return a photo of a goods
	router.HandleFunc("/goods/photo/{id:[0-9]+}", middleware.Chain(shared.GetGoodsPhoto(), middleware.Auth())).Methods("GET")

	// DELETE /goods/photo/{id:[0-9]+} will delete a photo of a goods
	router.HandleFunc("/goods/photo/{id:[0-9]+}", middleware.Chain(DeleteGoodsPhoto(), middleware.Auth())).Methods("DELETE")

	// GET /goods/{id:[0-9]+} will return a goods
	router.HandleFunc("/goods/id/{id:[0-9]+}", middleware.Chain(shared.GetGoodsByID(), middleware.Auth())).Methods("GET")

	// POST /goods/photos will upload photos of a goods
	router.HandleFunc("/goods/photos", middleware.Chain(UploadGoodsPhotos(), middleware.Auth())).Methods("POST")

	// POST /general_discount will create a new general discount
	router.HandleFunc("/general_discount", middleware.Chain(CreateGeneralDiscount(), middleware.Auth())).Methods("POST")

	// GET /general_discount/{menu_type:[0-9]+} will return a general discount
	router.HandleFunc("/general_discount/{menu_type:[0-9]+}", middleware.Chain(shared.GetGeneralDiscount(), middleware.Auth())).Methods("GET")

	// PATCH /general_discount will update a general discount
	router.HandleFunc("/general_discount", middleware.Chain(UpdateGeneralDiscount(), middleware.Auth())).Methods("PATCH")

	// DELETE /general_discount/{menu_type:[0-9]+} will delete a general discount
	router.HandleFunc("/general_discount/{menu_type:[0-9]+}", middleware.Chain(DeleteGeneralDiscount(), middleware.Auth())).Methods("DELETE")

	// POST /legend_information will create the legend information
	router.HandleFunc("/legend_information", middleware.Chain(CreateLegendInformation(), middleware.Auth())).Methods("POST")

	return router
}
