package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
)

func GetCTWInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctw_information, err := model.GetCTWInformation()
		if err != nil {
			log.Printf("%s: Error getting ctw information: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Error getting ctw information").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, ctw_information).Send(w)
	}
}
