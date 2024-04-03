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

func CreateGeneralDiscount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// we need menu_type, value, unit

		var general_discount model.GeneralDiscount

		err := json.NewDecoder(r.Body).Decode(&general_discount)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		// reject general_discount creation if at least one general_discount with same menu_type exists
		exists, err := model.CheckGeneralDiscountExistenceByMenuType(general_discount.MenuType)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}
		if exists {
			log.Printf("%s: %s", r.URL.Path, "General discount with same menu type exists")
			response.NewResponse("error", http.StatusBadRequest, "General discount with same menu type exists").Send(w)
			return
		}

		// create the general_discount
		err = model.CreateGeneralDiscount(&general_discount)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "General discount created").Send(w)
	}
}

func UpdateGeneralDiscount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// we need menu_type, value, unit

		var general_discount model.GeneralDiscount

		err := json.NewDecoder(r.Body).Decode(&general_discount)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		// update the general_discount
		err = model.UpdateGeneralDiscount(&general_discount)
		if err != nil {
			if err.Error() == "no fields to update" {
				log.Printf("%s: %s", r.URL.Path, err)
				response.NewResponse("error", http.StatusBadRequest, "No fields to update").Send(w)
				return
			}
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "General discount updated").Send(w)
	}
}

func DeleteGeneralDiscount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// we need menu_type

		menu_type_str := mux.Vars(r)["menu_type"]
		menu_type, err := strconv.Atoi(menu_type_str)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid menu_type").Send(w)
			return
		}

		// delete the general_discount
		err = model.DeleteGeneralDiscount(menu_type)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "General discount deleted").Send(w)
	}
}
