package admin

import (
	"Legend/model"
	"Legend/response"
	"encoding/json"
	"log"
	"net/http"
)

func CreateLegendInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var legend_information model.LegendInformation

		err := json.NewDecoder(r.Body).Decode(&legend_information)
		if err != nil {
			log.Printf("%s: Error decoding request body: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Error decoding request body").Send(w)
			return
		}

		err = model.CreateLegendInformation(&legend_information)
		if err != nil {
			if err.Error() == "legend information already exists" {
				log.Printf("%s: Legend information already exists", r.URL.Path)
				response.NewResponse("error", http.StatusConflict, "Legend information already exists").Send(w)
				return
			}
			log.Printf("%s: Error creating legend information: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Error creating legend information").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "Legend information created").Send(w)
	}
}

func UpdateLegendInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var legend_information model.LegendInformation

		err := json.NewDecoder(r.Body).Decode(&legend_information)
		if err != nil {
			log.Printf("%s: Error decoding request body: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Error decoding request body").Send(w)
			return
		}

		err = model.UpdateLegendInformation(&legend_information)
		if err != nil {
			if err.Error() == "no fields to update" {
				log.Printf("%s: No fields to update", r.URL.Path)
				response.NewResponse("error", http.StatusBadRequest, "No fields to update").Send(w)
				return
			}

			log.Printf("%s: Error updating legend information: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Error updating legend information").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Legend information updated").Send(w)
	}
}
