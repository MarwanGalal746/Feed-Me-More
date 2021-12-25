package handlers

import (
	"backend/pkg/domain/sandwich"
	"backend/pkg/driver"
	"backend/pkg/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	//this CORS to enable frontend request to the backend endpoints
	headers := handlers.AllowedHeaders([]string{"Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	dbConnection := driver.GetDbConnection()

	sandwichHandler := SandwichHandlers{service.NewSandwichService(sandwich.NewSandwichRepositoryDb(dbConnection))}

	//sandwiches endpoints
	router.HandleFunc("/api/sandwiches", sandwichHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/sandwiches", sandwichHandler.Read).Methods(http.MethodGet)
	router.HandleFunc("/api/sandwiches/{id:[0-9]+}", sandwichHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/sandwiches/{id:[0-9]+}", sandwichHandler.Delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))

}
