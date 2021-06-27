package db

import "log"

func MakeAdmin(rollno string) bool {
	stmt, _ := DB.Prepare("UPDATE User SET Admin=1 WHERE rollno=?")
	defer stmt.Close()
	_, err := stmt.Exec(rollno)
	if err != nil {
		log.Println("Cannot make admin.")
		return false
	}
	log.Printf("%v is now the admin.\n", rollno)
	return true
}

func RemoveAdmin(rollno string) bool {
	stmt, _ := DB.Prepare("UPDATE User SET Admin=0 WHERE rollno=?")
	defer stmt.Close()
	_, err := stmt.Exec(rollno)
	if err != nil {
		log.Println("Cannot remove admin.")
		return false
	}
	log.Printf("%v is no more the admin.\n", rollno)
	return true
}

func IsAdmin(rollno string) bool {
	var admin bool
	err := DB.QueryRow(`SELECT Admin FROM User WHERE rollno=$1`, rollno).Scan(&admin)
	if err != nil {
		log.Println("Cannot determine the status at the moment.")
		return false
	}
	return admin
}
