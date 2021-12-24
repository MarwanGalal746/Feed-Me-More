package handlers

import (
	"backend/pkg/domain/dessert"
	"backend/pkg/errs"
	"backend/pkg/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type DessertHandlers struct {
	service service.DessertService
}

func (dessertHandler DessertHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newDessert dessert.Dessert
	_ = json.NewDecoder(r.Body).Decode(&newDessert)
	id, err := dessertHandler.service.Create(newDessert)
	newDessert.ID = int(uint(id))
	//handling errors
	if err != nil && err.Error() == `ERROR: duplicate key value violates unique constraint "desserts_name_key" (SQLSTATE 23505)` {
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
	json.NewEncoder(w).Encode(newDessert)
}

func (dessertHandler DessertHandlers) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	desserts, err := dessertHandler.service.Read()
	//handling errors
	if err == sql.ErrNoRows || len(desserts) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoDessertsFound.Error(), http.StatusOK))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(desserts)
}

func (dessertHandler DessertHandlers) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	var dessert dessert.Dessert
	_ = json.NewDecoder(r.Body).Decode(&dessert)
	// validate inputs
	err := dessertHandler.service.Update(dessert, id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDessertNotFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errs.NewResponse("Dessert has been updated successfully", http.StatusOK))
}

func (dessertHandler DessertHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	err := dessertHandler.service.Delete(id)
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
	json.NewEncoder(w).Encode(errs.NewResponse("Dessert has been deleted successfully", http.StatusOK))
}
