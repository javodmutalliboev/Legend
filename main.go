//   Product Api:
//    version: 0.1
//    title: Product Api of Legend
//   Schemes: http, https
//   Host:
//   BasePath: /api/v1
//      Consumes:
//      - application/json
//   Produces:
//   - application/json
//   SecurityDefinitions:
//    Bearer:
//     type: apiKey
//     name: Authorization
//     in: header
//   swagger:meta
package main

import (
	"Legend/developer"
	"Legend/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := router.Router()

	go func() { developer.CreateAdminCLI() }()

	http.ListenAndServe(":8080", router)
}
