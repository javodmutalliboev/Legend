package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSocialNetworks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all social networks
		socialNetworks, err := model.GetSocialNetworks()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, socialNetworks).Send(w)
	}
}

func GetSocialNetworkIcon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the icon of a social network
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid ID").Send(w)
			return
		}

		iconPtr, err := model.GetSocialNetworkIcon(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		ContentType := http.DetectContentType(*iconPtr)
		w.Header().Set("Content-Type", ContentType)
		w.Write(*iconPtr)
	}
}
