package main

//import (
//	"database/sql"
//	"fmt"
//)
//
//type MealRepositoryDb struct {
//	db *sql.DB
//}
//
//func (r MealRepositoryDb) Create(newMeal Meal) (int, error) {
//	row := r.db.QueryRow(`INSERT INTO meal (name) VALUES ($1) RETURNING id`, newMeal.Name)
//	var id int
//	err := row.Scan(&id)
//	fmt.Println(id)
//	return id, err
//}
//
//func (r MealRepositoryDb) Read() ([]Meal, error) {
//	meals := make([]Meal, 0)
//	rows, err := r.db.Query("SELECT id, name FROM meal")
//	if err != nil {
//		return meals, err
//	}
//	for rows.Next() {
//		var meal Meal
//		switch err := rows.Scan(&meal.ID, &meal.Name); err {
//		case sql.ErrNoRows:
//			return meals, sql.ErrNoRows
//		case nil:
//			meals = append(meals, meal)
//		default:
//			return meals, ErrDb
//		}
//	}
//	return meals, nil
//}
//
//func (r MealRepositoryDb) Update(meal Meal, id string) error {
//	var name string
//	row := r.db.QueryRow("SELECT name FROM meal WHERE id= $1", id)
//	err := row.Scan(&name)
//	if err != nil {
//		return err
//	}
//	_, err = r.db.Query("UPDATE meal SET name= $1 WHERE id=$2", meal.Name, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (r MealRepositoryDb) Delete(id string) error {
//	var name string
//	row := r.db.QueryRow("SELECT name FROM meal WHERE id=$1", id)
//	err := row.Scan(&name)
//	if err != nil {
//		return err
//	}
//	_, err = r.db.Query("DELETE FROM meal WHERE id= $1", id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func NewMealRepositoryDb(db *sql.DB) MealRepositoryDb {
//	return MealRepositoryDb{db}
//}
