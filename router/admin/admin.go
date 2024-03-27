package admin

import (
	"Legend/middleware"
	"Legend/response"

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
	router.HandleFunc("/menus/{type:[0-9]+}", middleware.Chain(GetMenus(), middleware.Auth())).Methods("GET")

	// PATCH /menu will update a menu
	router.HandleFunc("/menu", middleware.Chain(UpdateMenu(), middleware.Auth())).Methods("PATCH")

	// DELETE /menu/{id:[0-9]+} will delete a menu
	router.HandleFunc("/menu/{id:[0-9]+}", middleware.Chain(DeleteMenu(), middleware.Auth())).Methods("DELETE")

	// GET /menu/types will return all menu types
	router.HandleFunc("/menu/types", middleware.Chain(GetMenuTypes(), middleware.Auth())).Methods("GET")

	// POST /goods will create a new goods
	router.HandleFunc("/goods", middleware.Chain(CreateGoods(), middleware.Auth())).Methods("POST")

	return router
}
