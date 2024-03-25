package admin

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
)

func GetMenuTypes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		MenuTypes, err := model.GetMenuTypes()

		if err != nil {
			log.Printf("%s: Error getting menu types: %v", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, MenuTypes).Send(w)
	}
}
