package session

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   []byte
	store *sessions.CookieStore
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize key and store after loading .env file
	key = []byte(os.Getenv("SESSION_KEY"))
	store = sessions.NewCookieStore(key)
}

func Session(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "session")

	return session
}

func SaveOptions(session *sessions.Session, MaxAge int /* in seconds */) {
	session.Options = &sessions.Options{
		Path: "/",
		MaxAge:/* 1 day */ MaxAge,
		HttpOnly: true,
		Secure:   false,
		Domain:   os.Getenv("ADMIN_ORIGIN_COOKIE"),
		SameSite: http.SameSiteLaxMode,
	}
}
