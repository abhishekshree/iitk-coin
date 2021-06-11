package db

import (
	"database/sql"
	"fmt"
)

type User struct {
	Rollno   string
	Name     string
	Password string
}

var DB *sql.DB

func Add(user User) bool {
	status := false

	checker, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, user.Rollno)

	if checker.Next() {
		fmt.Printf("%v with roll number %v already EXISTS.\n", user.Name, user.Rollno)
		return status
	}

	statement, err := DB.Prepare("INSERT INTO User (rollno, name, password) VALUES (?, ?, ?)")

	if err == nil {
		statement.Exec(user.Rollno, user.Name, user.Password)
		fmt.Printf("Added %v with roll number %v.\n", user.Name, user.Rollno)
		status = true
	} else {
		panic(err)
	}

	return status
}

func FindPass(rollno string) (string, bool) {
	row, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, rollno)
	user := User{}
	if row.Next() {
		err := row.Scan(&user.Rollno, &user.Name, &user.Password)
		if err != nil {
			return "User does not exists", false
		}
		return user.Password, true
	}
	return "", false
}

func UserExists(rollno string) bool {
	checker, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, rollno)
	return checker.Next()
}

func Clean() bool {
	stmt, err := DB.Prepare("DROP TABLE IF EXISTS User")
	if err == nil {
		stmt.Exec()
		return true
	}
	return false
}
