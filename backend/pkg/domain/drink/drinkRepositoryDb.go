package drink

import (
	"backend/pkg/errs"
	"database/sql"
	"fmt"
)

type DrinkRepositoryDb struct {
	db *sql.DB
}

func (r DrinkRepositoryDb) Create(newDrink Drink) (int, error) {
	row := r.db.QueryRow(`INSERT INTO drink (name) VALUES ($1) RETURNING id`, newDrink.Name)
	var id int
	err := row.Scan(&id)
	fmt.Println(id)
	return id, err
}

func (r DrinkRepositoryDb) Read() ([]Drink, error) {
	drinks := make([]Drink, 0)
	rows, err := r.db.Query("SELECT id, name FROM drink")
	if err != nil {
		return drinks, err
	}
	for rows.Next() {
		var drink Drink
		switch err := rows.Scan(&drink.ID, &drink.Name); err {
		case sql.ErrNoRows:
			return drinks, sql.ErrNoRows
		case nil:
			drinks = append(drinks, drink)
		default:
			return drinks, errs.ErrDb
		}
	}
	return drinks, nil
}

func (r DrinkRepositoryDb) Update(drink Drink, id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM drink WHERE id= $1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("UPDATE drink SET name= $1 WHERE id=$2", drink.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r DrinkRepositoryDb) Delete(id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM drink WHERE id=$1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("DELETE FROM drink WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewDrinkRepositoryDb(db *sql.DB) DrinkRepositoryDb {
	return DrinkRepositoryDb{db}
}
