package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abhishekshree/iitk-coin/config"
)

type User struct {
	Rollno   string
	Name     string
	Password string
	Coins    int
}

var DB *sql.DB

func Add(user User) bool {
	status := false

	checker, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, user.Rollno)
	defer checker.Close()
	if checker.Next() {
		fmt.Printf("%v with roll number %v already EXISTS.\n", user.Name, user.Rollno)
		return status
	}

	statement, err := DB.Prepare("INSERT INTO User (rollno, name, password, coins) VALUES (?, ?, ?, ?)")

	if err == nil {
		statement.Exec(user.Rollno, user.Name, user.Password, config.INITIAL_COINS)
		fmt.Printf("Added %v with roll number %v.\n", user.Name, user.Rollno)
		status = true
	} else {
		panic(err)
	}

	return status
}

func FindPass(rollno string) (string, bool) {
	var pass string
	err := DB.QueryRow(`SELECT password FROM User WHERE rollno=$1`, rollno).Scan(&pass)
	if err != nil {
		return "User does not exists", false
	}
	return pass, true
}

func UserExists(rollno string) bool {
	checker, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, rollno)
	defer checker.Close()
	return checker.Next()
}

func Clean() bool {
	stmt, _ := DB.Prepare("DROP TABLE IF EXISTS Transactions")
	_, err := stmt.Exec()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
