package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configure the database connection (always check errors)
	// db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	db, _ := sql.Open("mysql", "7N10fxmDYp:NnktEISo8u@(remotemysql.com:3306)/7N10fxmDYp?parseTime=true")

	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
	);`

	_, _ = db.Exec(query)

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	_, _ = db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
	defer rows.Close()

	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
		users = append(users, u)
	}
	err := rows.Err() // check err

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	//err := db.Ping()

	insert, _ := db.Query("INSERT INTO employees VALUES ( 'SampleData', 391 )")

	fmt.Println(insert)
}

// https://gowebexamples.com/mysql-database/
