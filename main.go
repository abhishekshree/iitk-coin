package main

import (
	"database/sql"
	"log"

	"github.com/abhishekshree/iitk-coin/db"
	"github.com/abhishekshree/iitk-coin/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	DB, _ := sql.Open("sqlite3", "./Users.db")
	db.DB = DB
	defer DB.Close()

	// Creates the table User if it does not exists already.
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS User (rollno TEXT PRIMARY KEY, name TEXT, password TEXT, coins REAL, Admin BOOLEAN DEFAULT 0)")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'User' is Ready!")

	// Creates the table Transactions if it does not exists already.
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS Transactions (id INTEGER PRIMARY KEY, from_roll TEXT, to_roll TEXT, type TEXT, timestamp TEXT,amount_before_tax REAL, tax REAL)")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table 'Transactions' is Ready!")

	// App setup
	app := fiber.New()
	app.Use(logger.New())

	// Routes
	app.Get("/", routes.Hello)
	app.Post("/signup", routes.Signup)
	app.Post("/login", routes.Login)

	app.Get("/getCoins", routes.GetCoins)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// This will need a JWT bearer token.
	app.Get("/secretpage", routes.Secret)
	app.Post("/awardCoins", routes.AwardCoins)
	app.Post("/transferCoins", routes.TransferCoins)

	log.Fatal(app.Listen(":3000"))
}
