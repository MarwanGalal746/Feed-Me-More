package handlers

import (
	"backend/pkg/domain/sandwich"
	"backend/pkg/errs"
	"backend/pkg/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type SandwichHandlers struct {
	service service.SandwichService
}

func (sandwichHandler SandwichHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newSandwich sandwich.Sandwich
	_ = json.NewDecoder(r.Body).Decode(&newSandwich)
	id, err := sandwichHandler.service.Create(newSandwich)
	newSandwich.ID = int(uint(id))
	//handling errors
	if err != nil && err.Error() == `ERROR: duplicate key value violates unique constraint "sandwiches_name_key" (SQLSTATE 23505)` {
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
	json.NewEncoder(w).Encode(newSandwich)
}

func (sandwichHandler SandwichHandlers) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	sandwiches, err := sandwichHandler.service.Read()
	//handling errors
	if err == sql.ErrNoRows || len(sandwiches) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrNoSandwichesFound.Error(), http.StatusOK))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sandwiches)
}

func (sandwichHandler SandwichHandlers) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	var sandwich sandwich.Sandwich
	_ = json.NewDecoder(r.Body).Decode(&sandwich)
	// validate inputs
	err := sandwichHandler.service.Update(sandwich, id)
	//handling errors
	if err != nil {
		if err.Error() == `sql: no rows in result set` {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrSandwichNotFound.Error(), http.StatusBadRequest))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		}
		return
	}
	//sending the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errs.NewResponse("Sandwich has been updated successfully", http.StatusOK))
}

func (sandwichHandler SandwichHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	id := params["id"]
	err := sandwichHandler.service.Delete(id)
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
	json.NewEncoder(w).Encode(errs.NewResponse("Sandwich has been deleted successfully", http.StatusOK))
}
