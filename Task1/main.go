package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./Users.db")

	// To do a fresh run, uncomment the line below
	clean(db)

	// Creates the table if it does not exists already.
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS User (rollno TEXT PRIMARY KEY, name TEXT)")
	statement.Exec()
	fmt.Println("Table 'User' is Ready!")

	// Adding some elements
	usrs := []struct {
		rollno string
		name   string
	}{
		{"200028", "Abhishek Shree"},
		{"200029", "Someone"},
		{"200030", "Someone Else"},
		{"180199", "Bhuvan Singla"},
	}

	for _, usr := range usrs {
		add(db, usr.rollno, usr.name)
	}

	db.Close()
	fmt.Println("Done!")
}
