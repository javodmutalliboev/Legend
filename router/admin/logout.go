package admin

import (
	"Legend/response"
	"Legend/session"
	"log"
	"net/http"
)

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ses := session.Session(r)
		ses.Options.MaxAge = -1
		err := ses.Save(r, w)
		if err != nil {
			log.Printf("%s: Error saving session: %v", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Logged out").Send(w)
	}
}
