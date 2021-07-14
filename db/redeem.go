package db

import (
	"log"
	"time"

	"github.com/abhishekshree/iitk-coin/config"
)

func AddRedeemRequest(roll string, item string) bool {
	if _, ok := config.REDEEM_LIST[item]; !ok {
		return false
	}

	stmt, err := DB.Prepare("INSERT INTO RedeemRequests (rollno, item, timestamp) VALUES (?,?,?)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(roll, item, time.Now().Format(time.RFC850))
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		log.Println("Failed to add redeem request.")
		return false
	}
	return true
}

func updateRedeemStatus(id int, status int) bool {
	stmt, err := DB.Prepare("UPDATE RedeemRequests SET status=? WHERE id=?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(status, id)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		log.Println("Failed to update redeem status.")
		return false
	}
	return true
}

func RejectRedeemRequest(id int) bool {
	return updateRedeemStatus(id, 2)
}

// get rollno and item from RedeemRequests table
func GetRedeemRequest(id int) (roll string, item string, status int) {
	stmt, err := DB.Prepare("SELECT rollno, item, status FROM RedeemRequests WHERE id=?")
	if err != nil {
		log.Println(err)
	}
	err = stmt.QueryRow(id).Scan(&roll, &item, &status)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		log.Println("Failed to get redeem request.")
		return "", "", -1
	}
	return roll, item, status
}

func AcceptRedeemRequest(id int) bool {
	tx, err := DB.Begin()
	if err != nil {
		log.Println(err)
	}

	rollno, item, status := GetRedeemRequest(id)
	if status == 1 {
		return true
	}
	if status == 2 {
		return false
	}
	amt := config.REDEEM_LIST[item]

	stmt, _ := tx.Prepare("UPDATE User SET coins=coins-? WHERE rollno=?")
	_, err = stmt.Exec(amt, rollno)
	defer stmt.Close()
	if err != nil {
		log.Println("Cannot redeem, failed to update coins.")
		tx.Rollback()
		return false
	}

	stmt2, err := tx.Prepare("UPDATE RedeemRequests SET status=? WHERE id=?")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt2.Exec(1, id)
	defer stmt2.Close()
	if err != nil {
		log.Println(err)
		log.Println("Failed to update redeem status.")
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

func RejectPendingRedeemRequests(roll string) bool {
	tx, err := DB.Begin()
	if err != nil {
		log.Println(err)
	}
	stmt, err := tx.Prepare("UPDATE RedeemRequests SET status=? WHERE (rollno=? AND status=0)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(2, roll)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		log.Println("Failed to update redeem status.")
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
