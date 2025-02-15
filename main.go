package main

import (
	"Legend/developer"
	"Legend/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	router := router.Router()

	go func() { developer.CreateAdminCLI() }()

	LegendPort := os.Getenv("LEGEND_PORT")

	http.ListenAndServe(fmt.Sprintf(":%s", LegendPort), router)
}
