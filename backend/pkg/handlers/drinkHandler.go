package handlers

import (
	"backend/pkg/domain/drink"
	"backend/pkg/errs"
	"backend/pkg/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DrinkHandlers struct {
	service service.DrinkService
}

func (drinkHandler DrinkHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newDrink drink.Drink
	_ = json.NewDecoder(r.Body).Decode(&newDrink)
	id, err := drinkHandler.service.Create(newDrink)
	newDrink.ID = int(uint(id))
	//handling errors
	if err != nil && err.Error() == `ERROR: duplicate key value violates unique constraint "drinks_name_key" (SQLSTATE 23505)` {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDuplicateValue.Error(), http.StatusBadRequest))
		return
	} else if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newDrink)
}

func (drinkHandler DrinkHandlers) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	drinks, err := drinkHandler.service.Read()
	//handling errors
	if err == sql.ErrNoRows || len(drinks) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoDrinksFound.Error(), http.StatusOK))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drinks)
}

func (drinkHandler DrinkHandlers) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	var drink drink.Drink
	_ = json.NewDecoder(r.Body).Decode(&drink)
	// validate inputs
	err := drinkHandler.service.Update(drink, id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDrinkNotFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errs.NewResponse("Drink has been updated successfully", http.StatusOK))
}

func (drinkHandler DrinkHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	err := drinkHandler.service.Delete(id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoRowsFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errs.NewResponse("Drink has been deleted successfully", http.StatusOK))
}
