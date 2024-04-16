package admin

import (
	"Legend/interface_package"
	"Legend/model"
	"Legend/response"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreatePaymentMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var paymentMethod interface_package.PaymentMethod = &model.PaymentMethod{}

		if r.FormValue("name") == "" {
			log.Printf("%s: %s", r.URL.Path, "name is required")
			response.NewResponse("error", http.StatusBadRequest, "name is required").Send(w)
			return
		}

		_, logo_header, err := r.FormFile("logo")
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "logo is required").Send(w)
			return
		}

		if !strings.HasPrefix(logo_header.Header.Get("Content-Type"), "image/") {
			log.Printf("%s: %s", r.URL.Path, "logo file is not an image")
			response.NewResponse("error", http.StatusBadRequest, "logo file is not an image").Send(w)
			return
		}

		if logo_header.Size > 1024*1024 { // 1MB
			log.Printf("%s: %s", r.URL.Path, "logo file size is larger than 1MB")
			response.NewResponse("error", http.StatusBadRequest, "logo file size is larger than 1MB").Send(w)
			return
		}

		var logo []byte
		{
			f, err := logo_header.Open()
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err)
				response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
				return
			}
			defer f.Close()

			logo, err = io.ReadAll(f)
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err)
				response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
				return
			}
		}

		paymentMethod = &model.PaymentMethod{
			Name: r.FormValue("name"),
			Logo: logo,
		}

		err = paymentMethod.Create()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "A new payment method created").Send(w)
	}
}

func GetPaymentMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid ID").Send(w)
			return
		}

		paymentMethod, err := model.GetPaymentMethod(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, paymentMethod).Send(w)
	}
}

func UpdatePaymentMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var paymentMethod interface_package.PaymentMethod = &model.PaymentMethod{}

		if r.FormValue("id") == "" {
			log.Printf("%s: %s", r.URL.Path, "id is required")
			response.NewResponse("error", http.StatusBadRequest, "id is required").Send(w)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid ID").Send(w)
			return
		}

		paymentMethod = &model.PaymentMethod{
			ID: id,
		}

		if r.FormValue("name") != "" {
			paymentMethod.(*model.PaymentMethod).Name = r.FormValue("name")
		}

		_, logo_header, err := r.FormFile("logo")
		if err == nil {
			if !strings.HasPrefix(logo_header.Header.Get("Content-Type"), "image/") {
				log.Printf("%s: %s", r.URL.Path, "logo file is not an image")
				response.NewResponse("error", http.StatusBadRequest, "logo file is not an image").Send(w)
				return
			}

			if logo_header.Size > 1024*1024 { // 1MB
				log.Printf("%s: %s", r.URL.Path, "logo file size is larger than 1MB")
				response.NewResponse("error", http.StatusBadRequest, "logo file size is larger than 1MB").Send(w)
				return
			}

			var logo []byte
			{
				f, err := logo_header.Open()
				if err != nil {
					log.Printf("%s: %s", r.URL.Path, err)
					response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
					return
				}
				defer f.Close()

				logo, err = io.ReadAll(f)
				if err != nil {
					log.Printf("%s: %s", r.URL.Path, err)
					response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
					return
				}
			}

			paymentMethod.(*model.PaymentMethod).Logo = logo
		}

		err = paymentMethod.Update()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Payment method updated").Send(w)
	}
}

func DeletePaymentMethod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid ID").Send(w)
			return
		}

		err = model.DeletePaymentMethod(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Payment method deleted").Send(w)
	}
}
