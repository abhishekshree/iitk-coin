package db

import (
	"log"
	"time"
)

func AddAwardLog(from string, to string, amount float64) bool {
	// Tax column would be zero here.
	stmt, err := DB.Prepare("INSERT INTO Transactions (from_roll, to_roll, type, timestamp, amount_before_tax, tax) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(from, to, "Awarded", time.Now().Format(time.RFC850), amount, 0)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		log.Println("Failed to add data to award log")
		return false
	}
	return true
}

func AddTransferLog(from string, to string, amount float64, tax float64) bool {
	stmt, _ := DB.Prepare("INSERT INTO Transactions (from_roll, to_roll, type, timestamp, amount_before_tax, tax) VALUES (?,?,?,?,?,?)")
	_, err := stmt.Exec(from, to, "Transferred", time.Now().Format(time.RFC850), amount, tax)
	defer stmt.Close()
	if err != nil {
		log.Println("Failed to add data to award log")
		return false
	}
	return true
}
