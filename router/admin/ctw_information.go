package admin

import (
	"Legend/model"
	"Legend/response"
	"encoding/json"
	"log"
	"net/http"
)

func CreateCTWInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctw_information model.CTWInformation

		err := json.NewDecoder(r.Body).Decode(&ctw_information)
		if err != nil {
			log.Printf("%s: Error decoding request body: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Error decoding request body").Send(w)
			return
		}

		err = model.CreateCTWInformation(&ctw_information)
		if err != nil {
			if err.Error() == "ctw information already exists" {
				log.Printf("%s: CTW information already exists", r.URL.Path)
				response.NewResponse("error", http.StatusConflict, "CTW information already exists").Send(w)
				return
			}
			log.Printf("%s: Error creating ctw information: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Error creating ctw information").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "CTW information created").Send(w)
	}
}
