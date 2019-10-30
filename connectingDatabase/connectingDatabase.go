package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configure the database connection (always check errors)
	// db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	db, _ := sql.Open("mysql", "7N10fxmDYp:NnktEISo8u@(remotemysql.com:3306)/7N10fxmDYp?parseTime=true")

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	//err := db.Ping()

	insert, _ := db.Query("INSERT INTO employees VALUES ( 'SampleData', 391 )")

	fmt.Println(insert)
}
