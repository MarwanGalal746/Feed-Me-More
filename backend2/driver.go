package main

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//	"log"
//)
//
//func GetDbConnection() *sql.DB {
//	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
//		"localhost",
//		"5432",
//		"postgres",
//		"fmm-user",
//		"fmm-password")
//	sqlDB, err := sql.Open("postgres", dataSourceName)
//	if err != nil {
//		fmt.Println(err)
//		log.Fatal(fmt.Sprintf("unable to conect to db"))
//		panic(err)
//	}
//	log.Println("connected to db ")
//	log.Println("pinged db")
//	return sqlDB
//}
