package middleware

import (
	"Legend/response"
	"Legend/session"
	"net/http"
)

func Auth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ses := session.Session(r)

			// Check if user is authenticated
			if auth, ok := ses.Values["authenticated"].(bool); !ok || !auth {
				response.NewResponse("error", http.StatusForbidden, "Forbidden").Send(w)
				return
			}

			session.SaveOptions(ses, 24*60*60)
			ses.Save(r, w)

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
