package main

import (
	"Legend/developer"
	"Legend/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := router.Router()

	go func() { developer.CreateAdminCLI() }()

	http.ListenAndServe(":8080", router)
}
