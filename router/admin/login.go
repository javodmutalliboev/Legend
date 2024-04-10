package admin

import (
	"Legend/database"
	"Legend/model"
	"Legend/password"
	"Legend/response"
	"Legend/session"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var admin model.Admin
		err := json.NewDecoder(r.Body).Decode(&admin)
		if err != nil {
			log.Printf("%s: Error decoding json: %v", r.URL.Path, err)
			response.NewResponse("error", http.StatusBadRequest, "Invalid request").Send(w)
			return
		}

		// email
		if admin.Email == "" {
			response.NewResponse("error", http.StatusBadRequest, "Email is required").Send(w)
			return
		}

		// password
		if admin.Password == "" {
			response.NewResponse("error", http.StatusBadRequest, "Password is required").Send(w)
			return
		}

		// authenticate
		authenticated, err := authenticate(admin.Email, admin.Password)
		if err != nil {
			log.Printf("%s: Error authenticating: %v", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		if !authenticated {
			response.NewResponse("error", http.StatusUnauthorized, "Invalid credentials").Send(w)
			return
		}

		ses := session.Session(r)
		ses.Values["authenticated"] = true
		session.SaveOptions(ses, 24*60*60)

		err = ses.Save(r, w)
		if err != nil {
			log.Printf("%s: Error saving session: %v", r.URL.Path, err)
			response.NewResponse("error", http.StatusInternalServerError, "Internal server error").Send(w)
			return
		}

		response.NewResponse("success", http.StatusOK, "Authenticated").Send(w)
	}
}

// authenticate function
func authenticate(email, pass string) (bool, error) {
	// connect to database
	database := database.DB()
	defer database.Close()

	// query
	row := database.QueryRow("SELECT id, name, surname, email, password, created_at, updated_at FROM admin WHERE email = $1", email)

	// admin
	var admin model.Admin
	err := row.Scan(&admin.ID, &admin.Name, &admin.Surname, &admin.Email, &admin.Password, &admin.CreatedAt, &admin.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // or return false, errors.New("Admin not found")
		}
		return false, err
	}

	// compare password
	authenticated := password.CheckPasswordHash(pass, admin.Password)
	if !authenticated {
		return false, nil
	}

	return true, nil
}
