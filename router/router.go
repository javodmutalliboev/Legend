package router

import (
	"Legend/middleware"
	"Legend/response"
	"Legend/router/admin"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(response.Response{Status: "success", Code: http.StatusOK, Data: "legend server"})
	}).Methods("GET")

	// attach admin.AdminRouter() to /api/v1 so that the path will be /api/v1/admin
	router.PathPrefix("/admin").Handler(admin.AdminRouter())

	// Apply the CORS middleware to the router
	router.Use(middleware.CORSMiddleware)

	return router
}
