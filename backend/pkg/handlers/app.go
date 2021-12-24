package handlers

import (
	"backend/pkg/domain/dessert"
	"backend/pkg/domain/drink"
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

	drinkHandler := DrinkHandlers{service.NewDrinkService(drink.NewDrinkRepositoryDb(dbConnection))}
	dessertHandler := DessertHandlers{service.NewDessertService(dessert.NewDessertRepositoryDb(dbConnection))}
	sandwichHandler := DessertHandlers{service.NewDessertService(dessert.NewDessertRepositoryDb(dbConnection))}

	//drinks endpoints
	router.HandleFunc("/api/drinks", drinkHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/drinks", drinkHandler.Read).Methods(http.MethodGet)
	router.HandleFunc("/api/drinks/{id:[0-9]+}", drinkHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/drinks/{id:[0-9]+}", drinkHandler.Delete).Methods(http.MethodDelete)

	//Desserts endpoints
	router.HandleFunc("/api/desserts", dessertHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/desserts", dessertHandler.Read).Methods(http.MethodGet)
	router.HandleFunc("/api/desserts/{id:[0-9]+}", dessertHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/desserts/{id:[0-9]+}", dessertHandler.Delete).Methods(http.MethodDelete)

	//sandwiches endpoints
	router.HandleFunc("/api/sandwiches", sandwichHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/sandwiches", sandwichHandler.Read).Methods(http.MethodGet)
	router.HandleFunc("/api/sandwiches/{id:[0-9]+}", sandwichHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/sandwiches/{id:[0-9]+}", sandwichHandler.Delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS(headers, methods, origins)(router)))

}
