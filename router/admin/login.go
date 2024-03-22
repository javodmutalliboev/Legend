package admin

import (
	"log"
	"net/http"
)

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: someone is hitting me", r.URL.Path)
	}
}
