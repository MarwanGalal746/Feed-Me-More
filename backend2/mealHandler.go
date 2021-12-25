package main

//type MealHandlers struct {
//	service MealService
//}
//
//func (mealHandler MealHandlers) Create(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("Content-Type", "application/json")
//	var newMeal Meal
//	_ = json.NewDecoder(r.Body).Decode(&newMeal)
//	id, err := mealHandler.service.Create(newMeal)
//	newMeal.ID = int(uint(id))
//	//handling errors
//	if err != nil && err.Error() == `ERROR: duplicate key value violates unique constraint "meals_name_key" (SQLSTATE 23505)` {
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(NewResponse(ErrDuplicateValue.Error(), http.StatusBadRequest))
//		return
//	} else if err != nil {
//		fmt.Println(err)
//		w.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
//		return
//	}
//	//sending the response
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(newMeal)
//}
//
//func (mealHandler MealHandlers) Read(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("Content-Type", "application/json")
//	meales, err := mealHandler.service.Read()
//	//handling errors
//	if err == sql.ErrNoRows || len(meales) == 0 {
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(NewResponse(ErrNoMealsFound.Error(), http.StatusOK))
//		return
//	} else if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
//		return
//	}
//	//sending the response
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(meales)
//}
//
//func (mealHandler MealHandlers) Update(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("Content-Type", "application/json")
//	params := mux.Vars(r) // Get params
//	id := params["id"]
//	var meal Meal
//	_ = json.NewDecoder(r.Body).Decode(&meal)
//	// validate inputs
//	err := mealHandler.service.Update(meal, id)
//	//handling errors
//	if err != nil {
//		if err.Error() == `sql: no rows in result set` {
//			w.WriteHeader(http.StatusBadRequest)
//			json.NewEncoder(w).Encode(NewResponse(ErrMealNotFound.Error(), http.StatusBadRequest))
//		} else {
//			w.WriteHeader(http.StatusInternalServerError)
//			json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
//		}
//		return
//	}
//	//sending the response
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(NewResponse("Meal has been updated successfully", http.StatusOK))
//}
//
//func (mealHandler MealHandlers) Delete(w http.ResponseWriter, r *http.Request) {
//	w.Header().Add("Content-Type", "application/json")
//	params := mux.Vars(r) // Get params
//	id := params["id"]
//	err := mealHandler.service.Delete(id)
//	//handling errors
//	if err != nil {
//		if err.Error() == `sql: no rows in result set` {
//			w.WriteHeader(http.StatusBadRequest)
//			json.NewEncoder(w).Encode(NewResponse(ErrNoRowsFound.Error(), http.StatusBadRequest))
//		} else {
//			w.WriteHeader(http.StatusInternalServerError)
//			json.NewEncoder(w).Encode(NewResponse(ErrDb.Error(), http.StatusInternalServerError))
//		}
//		return
//	}
//	//sending the response
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(NewResponse("Meal has been deleted successfully", http.StatusOK))
//}
