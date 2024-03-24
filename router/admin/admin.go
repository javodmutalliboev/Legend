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

	// POST /menu/{id}/sub will create a new sub menu
	router.HandleFunc("/menu/{id:[0-9]+}/sub", middleware.Chain(CreateSubMenu(), middleware.Auth())).Methods("POST")

	// POST /sub-menu/{id:[0-9]+}/sub will create a new sub menu
	router.HandleFunc("/sub-menu/{id:[0-9]+}/sub", middleware.Chain(CreateSubMenu(), middleware.Auth())).Methods("POST")

	return router
}
