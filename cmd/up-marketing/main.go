package main

import (
	"go-rest/router"
	"net/http"
)

func main() {
	r := router.NewRouter()

	// Start the server
	http.ListenAndServe(":8080", r)
}
