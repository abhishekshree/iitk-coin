package main

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type SignupInfo struct {
	Rollno   string `json:"Roll"`
	Username string `json:"Name"`
	Password string `json:"Password"`
}

type LoginInfo struct {
	Rollno   string `json:"Roll"`
	Password string `json:"Password"`
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Signup(c *fiber.Ctx) error {
	p := new(SignupInfo)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	user := User{
		p.Rollno,
		p.Username,
		hashAndSalt(p.Password),
	}
	status := add(user)

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": status,
		},
	)
}

func Login(c *fiber.Ctx) error {
	p := new(LoginInfo)
	if err := c.BodyParser(p); err != nil {
		return err
	}

	pass, status := findPass(p.Rollno)
	var res bool = false
	if status {
		res = ComparePasswords(pass, p.Password)
		if res {

			// Create token
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["roll"] = p.Rollno
			claims["loggedin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))

			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			return c.JSON(fiber.Map{"token": t, "status": res})
		}
	}
	return c.JSON(fiber.Map{"status": res})
}

func Secret(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	roll := claims["roll"].(string)
	loggedin := claims["loggedin"].(bool)

	// log.Println("The roll is ", roll)

	if loggedin && UserExists(roll) {
		return c.SendString("This is a very secret string.")
	}
	return c.SendString("Bad Token, user does not exist.")
}
