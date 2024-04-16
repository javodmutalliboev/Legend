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

func CreateSocialNetwork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var socialNetwork interface_package.SocialNetwork = &model.SocialNetwork{}

		if r.FormValue("name") == "" {
			log.Printf("%s: %s", r.URL.Path, "name is required")
			response.NewResponse("error", http.StatusBadRequest, "name is required").Send(w)
			return
		}

		if r.FormValue("url") == "" {
			log.Printf("%s: %s", r.URL.Path, "url is required")
			response.NewResponse("error", http.StatusBadRequest, "url is required").Send(w)
			return
		}

		_, icon_header, err := r.FormFile("icon")
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "icon is required").Send(w)
			return
		}

		if !strings.HasPrefix(icon_header.Header.Get("Content-Type"), "image/") {
			log.Printf("%s: %s", r.URL.Path, "icon file is not an image")
			response.NewResponse("error", http.StatusBadRequest, "icon file is not an image").Send(w)
			return
		}

		if icon_header.Size > 1024*1024 { // 1MB
			log.Printf("%s: %s", r.URL.Path, "icon file size is larger than 1MB")
			response.NewResponse("error", http.StatusBadRequest, "icon file size is larger than 1MB").Send(w)
			return
		}

		var icon []byte
		{
			f, err := icon_header.Open()
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err)
				response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
				return
			}
			defer f.Close()

			icon, err = io.ReadAll(f)
			if err != nil {
				log.Printf("%s: %s", r.URL.Path, err)
				response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
				return
			}
		}

		socialNetwork = &model.SocialNetwork{
			Name: r.FormValue("name"),
			Icon: icon,
			URL:  r.FormValue("url"),
		}

		err = socialNetwork.Create()
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusCreated, "A new social network created").Send(w)
	}
}

func GetSocialNetwork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid ID").Send(w)
			return
		}

		// Get a social network
		socialNetwork, err := model.GetSocialNetwork(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, socialNetwork).Send(w)
	}
}

func UpdateSocialNetwork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var socialNetwork interface_package.SocialNetwork = &model.SocialNetwork{}

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

		socialNetwork = &model.SocialNetwork{
			ID: id,
		}

		if r.FormValue("name") != "" {
			socialNetwork.(*model.SocialNetwork).Name = r.FormValue("name")
		}

		if r.FormValue("url") != "" {
			socialNetwork.(*model.SocialNetwork).URL = r.FormValue("url")
		}

		_, icon_header, err := r.FormFile("icon")
		if err == nil {
			if !strings.HasPrefix(icon_header.Header.Get("Content-Type"), "image/") {
				log.Printf("%s: %s", r.URL.Path, "icon file is not an image")
				response.NewResponse("error", http.StatusBadRequest, "icon file is not an image").Send(w)
				return
			}

			if icon_header.Size > 1024*1024 { // 1MB
				log.Printf("%s: %s", r.URL.Path, "icon file size is larger than 1MB")
				response.NewResponse("error", http.StatusBadRequest, "icon file size is larger than 1MB").Send(w)
				return
			}

			var icon []byte
			{
				f, err := icon_header.Open()
				if err != nil {
					log.Printf("%s: %s", r.URL.Path, err)
					response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
					return
				}
				defer f.Close()

				icon, err = io.ReadAll(f)
				if err != nil {
					log.Printf("%s: %s", r.URL.Path, err)
					response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
					return
				}
			}

			socialNetwork.(*model.SocialNetwork).Icon = icon
		}

		err = socialNetwork.Update()
		if err != nil {
			if err.Error() == "no values to update" {
				log.Printf("%s: %s", r.URL.Path, err)
				response.NewResponse("error", http.StatusBadRequest, "No values to update").Send(w)
				return
			}
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Social network updated").Send(w)
	}
}

func DeleteSocialNetwork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid ID").Send(w)
			return
		}

		err = model.DeleteSocialNetwork(id)
		if err != nil {
			log.Printf("%s: %s", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Social network deleted").Send(w)
	}
}
