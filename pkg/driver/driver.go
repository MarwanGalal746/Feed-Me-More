package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

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
	_, err = sqlDB.Exec(`DROP TABLE IF EXISTS sandwich;`)
	if err != nil {
		panic(err)
	}
	_, err = sqlDB.Exec(`CREATE TABLE sandwich (ID INT PRIMARY KEY NOT NULL, NAME text);`)
	if err != nil {
		panic(err)
	}
	return sqlDB
}
