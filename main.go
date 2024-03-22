package main

import (
	"Legend/router"
	"net/http"
)

func main() {
	router := router.Router()

	http.ListenAndServe(":8080", router)
}
