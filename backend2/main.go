package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	//handleRequests()

	Start()
}

type Meal struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MealRepository interface {
	Create(Meal) (int, error)
	Read() ([]Meal, error)
	Update(Meal, string) error
	Delete(string) error
}

type MealRepositoryDb struct {
	db *sql.DB
}

func (r MealRepositoryDb) Create(newMeal Meal) (int, error) {
	row := r.db.QueryRow(`INSERT INTO meal (id, name) VALUES ($1, $2) RETURNING id`, newMeal.ID, newMeal.Name)
	var id int
	err := row.Scan(&id)
	fmt.Println(r.db)
	fmt.Println(err)
	fmt.Println(id)
	return id, err
}

func (r MealRepositoryDb) Read() ([]Meal, error) {
	meals := make([]Meal, 0)
	rows, err := r.db.Query("SELECT id, name FROM meal")
	if err != nil {
		return meals, err
	}
	for rows.Next() {
		var meal Meal
		switch err := rows.Scan(&meal.ID, &meal.Name); err {
		case sql.ErrNoRows:
			return meals, sql.ErrNoRows
		case nil:
			meals = append(meals, meal)
		default:
			return meals, ErrDb
		}
	}
	return meals, nil
}

func (r MealRepositoryDb) Update(meal Meal, id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM meal WHERE id= $1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("UPDATE meal SET name= $1 WHERE id=$2", meal.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r MealRepositoryDb) Delete(id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM meal WHERE id=$1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("DELETE FROM meal WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewMealRepositoryDb(db *sql.DB) MealRepositoryDb {
	return MealRepositoryDb{db}
}

type MealService interface {
	Create(Meal) (int, error)
	Read() ([]Meal, error)
	Update(Meal, string) error
	Delete(string) error
}

type DefaultMealService struct {
	repo MealRepository
}

func (s DefaultMealService) Create(newMeal Meal) (int, error) {
	return s.repo.Create(newMeal)
}

func (s DefaultMealService) Read() ([]Meal, error) {
	return s.repo.Read()
}

func (s DefaultMealService) Update(targetedMeal Meal, id string) error {
	return s.repo.Update(targetedMeal, id)
}

func (s DefaultMealService) Delete(id string) error {
	return s.repo.Delete(id)
}

func NewMealService(repository MealRepository) DefaultMealService {
	return DefaultMealService{repo: repository}
}

type MealHandlers struct {
	service MealService
}

func (mealHandler MealHandlers) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	var newMeal Meal
	_ = json.NewDecoder(r.Body).Decode(&newMeal)
	fmt.Println(newMeal.Name)
	id, err := mealHandler.service.Create(newMeal)
	newMeal.ID = int(uint(id))
	//handling errors
	if err != nil && err.Error() == `ERROR: duplicate key value violates unique constraint "meals_name_key" (SQLSTATE 23505)` {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewResponse(ErrDuplicateValue.Error(), http.StatusBadRequest))
		return
	} else if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMeal)
}

func (mealHandler MealHandlers) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	meales, err := mealHandler.service.Read()
	//handling errors
	if err == sql.ErrNoRows || len(meales) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(NewResponse(ErrNoMealsFound.Error(), http.StatusOK))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meales)
}

func (mealHandler MealHandlers) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	var meal Meal
	_ = json.NewDecoder(r.Body).Decode(&meal)
	// validate inputs
	err := mealHandler.service.Update(meal, id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(NewResponse(ErrMealNotFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewResponse("Meal has been updated successfully", http.StatusOK))
}

func (mealHandler MealHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	err := mealHandler.service.Delete(id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(NewResponse(ErrNoRowsFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewResponse("Meal has been deleted successfully", http.StatusOK))
}

func GetDbConnection() *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "user", "mypassword", "user")
	sqlDB, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Println(err)
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")
	log.Println("pinged db")
	_, err = sqlDB.Exec(`DROP TABLE IF EXISTS meal;`)
	if err != nil {
		panic(err)
	}
	_, err = sqlDB.Exec(`CREATE TABLE meal (ID INT PRIMARY KEY NOT NULL, NAME text);`)
	if err != nil {
		panic(err)
	}

	return sqlDB
}

var (
	ErrDuplicateValue = errors.New("this value already exists")
	ErrDb             = errors.New("unexpected database error")
	ErrNoRowsFound    = errors.New("no values found")
	ErrMealNotFound   = errors.New("this sandwich is not found")
	ErrNoMealsFound   = errors.New("no desserts found")
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}

func Start() {

	//router := mux.NewRouter()
	////this CORS to enable frontend request to the backend endpoints
	//headers := handlers.AllowedHeaders([]string{"Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	//methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	//origins := handlers.AllowedOrigins([]string{"*"})
	//dbConnection := GetDbConnection()
	//
	//mealHandler := MealHandlers{NewMealService(NewMealRepositoryDb(dbConnection))}

	//meals endpoints
	http.HandleFunc("/", handle)
	//http.HandleFunc("/", mealHandler.Create)
	//http.HandleFunc("/read", mealHandler.Read)
	//http.HandleFunc("/update", mealHandler.Update)
	//http.HandleFunc("/delete", mealHandler.Delete)
	//router.HandleFunc("/api/meals", mealHandler.Create).Methods(http.MethodPost)
	//router.HandleFunc("/api/meals", mealHandler.Read).Methods(http.MethodGet)
	//router.HandleFunc("/api/meals/{id:[0-9]+}", mealHandler.Update).Methods(http.MethodPut)
	//router.HandleFunc("/api/meals/{id:[0-9]+}", mealHandler.Delete).Methods(http.MethodDelete)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handle(w http.ResponseWriter, r *http.Request) {
	dbConnection := GetDbConnection()
	mealHandler := MealHandlers{NewMealService(NewMealRepositoryDb(dbConnection))}
	switch r.Method {
	case http.MethodGet:
		mealHandler.Read(w, r)
	case http.MethodPost:
		mealHandler.Create(w, r)
	case http.MethodPut:
		mealHandler.Update(w, r)
	case http.MethodDelete:
		mealHandler.Delete(w, r)

	}
}
