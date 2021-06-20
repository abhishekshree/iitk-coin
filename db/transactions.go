package db

import (
	"log"
)

func CoinCount(rollno string) int {
	row, _ := DB.Query(`SELECT * FROM User WHERE rollno=$1`, rollno)
	defer row.Close()
	user := User{}
	if row.Next() {
		err := row.Scan(&user.Rollno, &user.Name, &user.Password, &user.Coins)
		if err != nil {
			return -1
		}
		return user.Coins
	}
	return -1
}

func AddCoins(rollno string, amt int) bool {
	if UserExists(rollno) {
		tx, err := DB.Begin()
		if err != nil {
			log.Println(err)
			return false
		}
		stmt, _ := tx.Prepare("UPDATE User SET coins=coins+? WHERE rollno=?")
		_, err = stmt.Exec(amt, rollno)
		if err != nil {
			log.Println("Cannot add coins.")
			tx.Rollback()
			return false
		}
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	return false
}

func TransferCoins(from string, to string, amt int) bool {
	if UserExists(from) && UserExists(to) {
		tx, err := DB.Begin()
		if err != nil {
			log.Println(err)
			return false
		}
		stmt, _ := tx.Prepare("UPDATE User SET coins=coins-? WHERE rollno=? AND coins - ? >= 0")
		_, err = stmt.Exec(amt, from, amt)
		if err != nil {
			log.Println("Cannot transfer coins.")
			tx.Rollback()
			return false
		}
		stmt, _ = tx.Prepare("UPDATE User SET coins=coins+? WHERE rollno=?")
		_, err = stmt.Exec(amt, to)
		if err != nil {
			log.Println("Cannot transfer coins.")
			tx.Rollback()
			return false
		}
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	return false
}
