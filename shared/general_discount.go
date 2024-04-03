package shared

import (
	"Legend/model"
	"Legend/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetGeneralDiscount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the menu_type
		menu_type_str := mux.Vars(r)["menu_type"]
		menu_type, err := strconv.Atoi(menu_type_str)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid menu_type").Send(w)
			return
		}

		// get the general discount
		general_discount, err := model.GetGeneralDiscountByMenuType(menu_type)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, general_discount).Send(w)
	}
}
