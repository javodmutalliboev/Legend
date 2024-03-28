package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMenus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		menuType, err := strconv.Atoi(vars["type"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		// Get all menus from the database
		menus, err := model.GetMenus(menuType)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, menus).Send(w)
	}
}
