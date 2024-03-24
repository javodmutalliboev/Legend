package admin

import (
	"Legend/model"
	"Legend/response"
	"encoding/json"
	"log"
	"net/http"
)

func CreateSubMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a new sub menu
		var subMenu model.Menu

		// Decode the incoming sub menu json
		err := json.NewDecoder(r.Body).Decode(&subMenu)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		if *subMenu.ParentID == 0 {
			log.Printf("%s: %s", r.URL.Path, "Parent ID is required")
			response.NewResponse("error", http.StatusBadRequest, "Parent ID is required").Send(w)
			return
		}

		if subMenu.Title == "" {
			log.Printf("%s: %s", r.URL.Path, "Title is required")
			response.NewResponse("error", http.StatusBadRequest, "Title is required").Send(w)
			return
		}

		// Save the sub menu to the database
		err = subMenu.SubMenuSave()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "Sub menu created").Send(w)
	}
}
