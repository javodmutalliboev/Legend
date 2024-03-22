package router

import (
	"Legend/response"
	"Legend/router/admin"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(response.Response{Status: "success", Code: http.StatusOK, Data: "legend server"})
	})

	adminRouter := admin.AdminRouter()
	router.PathPrefix("/admin").Handler(adminRouter)

	return router
}
