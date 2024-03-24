package admin

import (
	"Legend/model"
	"Legend/response"
	"encoding/json"
	"log"
	"net/http"
)

func CreateMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var menu model.Menu

		// Decode the incoming Menu json
		err := json.NewDecoder(r.Body).Decode(&menu)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		if menu.Title == "" {
			log.Printf("%s: %s", r.URL.Path, "Title is required")
			response.NewResponse("error", http.StatusBadRequest, "Title is required").Send(w)
			return
		}

		// Save the menu to the database
		err = menu.MenuSave()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "Menu created").Send(w)
	}
}

func GetMenus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all menus from the database
		menus, err := model.GetMenus()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, menus).Send(w)
	}
}
