package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func main() {
	DB, _ = sql.Open("sqlite3", "./Users.db")
	clean()
	// Creates the table if it does not exists already.
	statement, _ := DB.Prepare("CREATE TABLE IF NOT EXISTS User (rollno TEXT PRIMARY KEY, name TEXT, password TEXT)")
	statement.Exec()
	fmt.Println("Table 'User' is Ready!")

	// App setup
	app := fiber.New()
	app.Use(logger.New())

	// Routes
	app.Get("/", Hello)
	app.Post("/signup", Signup)
	app.Post("/login", Login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// This will need a JWT bearer token.
	app.Get("/secretpage", Secret)

	log.Fatal(app.Listen(":3000"))
}
