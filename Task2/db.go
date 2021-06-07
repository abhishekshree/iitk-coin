package main

import (
	"fmt"
)

type User struct {
	rollno   string
	name     string
	password string
}

func add(user User) bool {
	status := false

	checker, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, user.rollno)

	if checker.Next() {
		fmt.Printf("%v with roll number %v already EXISTS.\n", user.name, user.rollno)
		return status
	}

	statement, err := DB.Prepare("INSERT INTO User (rollno, name, password) VALUES (?, ?, ?)")

	if err == nil {
		statement.Exec(user.rollno, user.name, user.password)
		fmt.Printf("Added %v with roll number %v.\n", user.name, user.rollno)
		status = true
	} else {
		panic(err)
	}

	return status
}

func findPass(rollno string) (string, bool) {
	row, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, rollno)
	user := User{}
	if row.Next() {
		err := row.Scan(&user.rollno, &user.name, &user.password)
		if err != nil {
			return "User does not exists", false
		}
		return user.password, true
	}
	return "", false
}

func UserExists(rollno string) bool {
	checker, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, rollno)
	return checker.Next()
}

func clean() bool {
	stmt, err := DB.Prepare("DROP TABLE IF EXISTS User")
	if err == nil {
		stmt.Exec()
		return true
	}
	return false
}
