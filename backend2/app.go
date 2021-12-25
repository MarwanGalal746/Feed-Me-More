package main

//
//import (
//	"github.com/gorilla/handlers"
//	"github.com/gorilla/mux"
//	"log"
//	"net/http"
//)
//
//func Start() {
//	router := mux.NewRouter()
//	//this CORS to enable frontend request to the backend endpoints
//	headers := handlers.AllowedHeaders([]string{"Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
//	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
//	origins := handlers.AllowedOrigins([]string{"*"})
//	dbConnection := GetDbConnection()
//
//	mealHandler := MealHandlers{NewMealService(NewMealRepositoryDb(dbConnection))}
//
//	//meals endpoints
//	router.HandleFunc("/api/meals", mealHandler.Create).Methods(http.MethodPost)
//	router.HandleFunc("/api/meals", mealHandler.Read).Methods(http.MethodGet)
//	router.HandleFunc("/api/meals/{id:[0-9]+}", mealHandler.Update).Methods(http.MethodPut)
//	router.HandleFunc("/api/meals/{id:[0-9]+}", mealHandler.Delete).Methods(http.MethodDelete)
//
//	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS(headers, methods, origins)(router)))
//
//}
