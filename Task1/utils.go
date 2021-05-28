package main

import (
	"database/sql"
	"fmt"
)

func add(db *sql.DB, rollno int, name string) bool {
	status := false

	checker, _ := db.Query(`SELECT * FROM User WHERE rollno=$1`, rollno)

	if checker.Next() {
		fmt.Printf("%v with roll number %v already EXISTS.\n", name, rollno)
		return status
	}

	statement, err := db.Prepare("INSERT INTO User (rollno, name) VALUES (?, ?)")

	if err == nil {
		statement.Exec(rollno, name)
		fmt.Printf("Added %v with roll number %v.\n", name, rollno)
		status = true
	} else {
		panic(err)
	}

	return status
}

func clean(db *sql.DB) bool {
	stmt, err := db.Prepare("DROP TABLE IF EXISTS User")
	if err == nil {
		stmt.Exec()
		return true
	}
	return false
}
