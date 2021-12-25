package sandwich

import (
	"backend/pkg/errs"
	"database/sql"
	"fmt"
)

type SandwichRepositoryDb struct {
	db *sql.DB
}

func (r SandwichRepositoryDb) Create(newSandwich Sandwich) (int, error) {
	row := r.db.QueryRow(`INSERT INTO sandwich (id, name) VALUES ($1, $2) RETURNING id`,
		newSandwich.ID, newSandwich.Name)
	var id int
	err := row.Scan(&id)
	fmt.Println(id)
	return id, err
}

func (r SandwichRepositoryDb) Read() ([]Sandwich, error) {
	sandwichs := make([]Sandwich, 0)
	rows, err := r.db.Query("SELECT id, name FROM sandwich")
	if err != nil {
		return sandwichs, err
	}
	for rows.Next() {
		var sandwich Sandwich
		switch err := rows.Scan(&sandwich.ID, &sandwich.Name); err {
		case sql.ErrNoRows:
			return sandwichs, sql.ErrNoRows
		case nil:
			sandwichs = append(sandwichs, sandwich)
		default:
			return sandwichs, errs.ErrDb
		}
	}
	return sandwichs, nil
}

func (r SandwichRepositoryDb) Update(sandwich Sandwich, id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM sandwich WHERE id= $1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("UPDATE sandwich SET name= $1 WHERE id=$2", sandwich.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r SandwichRepositoryDb) Delete(id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM sandwich WHERE id=$1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("DELETE FROM sandwich WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewSandwichRepositoryDb(db *sql.DB) SandwichRepositoryDb {
	return SandwichRepositoryDb{db}
}
