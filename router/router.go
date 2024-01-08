package router

import (
	handler "go-rest/internal"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	// r.HandleFunc("/hello", handler.HelloWorldHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handler.GetUserHandler).Methods("GET")

	// Additional routes go here

	// Serve static files (if needed)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return r
}
