package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPaymentMethods() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get all payment methods
		paymentMethods, err := model.GetPaymentMethods()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, paymentMethods).Send(w)
	}
}

func GetPaymentMethodLogo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid ID").Send(w)
			return
		}

		// get the logo of a payment method
		logo, err := model.GetPaymentMethodLogo(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		ContentType := http.DetectContentType(*logo)
		w.Header().Set("Content-Type", ContentType)
		w.Write(*logo)
	}
}
