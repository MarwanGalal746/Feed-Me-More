package dessert

import (
	"backend/pkg/errs"
	"database/sql"
	"fmt"
)

type DessertRepositoryDb struct {
	db *sql.DB
}

func (r DessertRepositoryDb) Create(newDessert Dessert) (int, error) {
	row := r.db.QueryRow(`INSERT INTO dessert (name) VALUES ($1) RETURNING id`, newDessert.Name)
	var id int
	err := row.Scan(&id)
	fmt.Println(id)
	return id, err
}

func (r DessertRepositoryDb) Read() ([]Dessert, error) {
	desserts := make([]Dessert, 0)
	rows, err := r.db.Query("SELECT id, name FROM dessert")
	if err != nil {
		return desserts, err
	}
	for rows.Next() {
		var dessert Dessert
		switch err := rows.Scan(&dessert.ID, &dessert.Name); err {
		case sql.ErrNoRows:
			return desserts, sql.ErrNoRows
		case nil:
			desserts = append(desserts, dessert)
		default:
			return desserts, errs.ErrDb
		}
	}
	return desserts, nil
}

func (r DessertRepositoryDb) Update(dessert Dessert, id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM dessert WHERE id= $1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("UPDATE dessert SET name= $1 WHERE id=$2", dessert.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r DessertRepositoryDb) Delete(id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM dessert WHERE id=$1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("DELETE FROM dessert WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewDessertRepositoryDb(db *sql.DB) DessertRepositoryDb {
	return DessertRepositoryDb{db}
}
