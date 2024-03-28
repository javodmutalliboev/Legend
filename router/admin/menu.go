package admin

import (
	"Legend/model"
	"Legend/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

		if menu.Type == 0 {
			log.Printf("%s: %s", r.URL.Path, "Type is required")
			response.NewResponse("error", http.StatusBadRequest, "Type is required").Send(w)
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

func UpdateMenu() http.HandlerFunc {
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

		// Update the menu in the database
		err = menu.MenuUpdate()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Menu updated").Send(w)
	}
}

func DeleteMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		// Delete the menu from the database
		err = model.DeleteMenu(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Menu deleted").Send(w)
	}
}
