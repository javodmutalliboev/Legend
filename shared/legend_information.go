package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
)

func GetLegendInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		legend_information, err := model.GetLegendInformation()
		if err != nil {
			log.Printf("%s: Error getting legend_information: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Error getting legend_information").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, legend_information).Send(w)
	}
}
