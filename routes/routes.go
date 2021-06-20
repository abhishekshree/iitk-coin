package routes

import (
	"time"

	"github.com/abhishekshree/iitk-coin/db"
	util "github.com/abhishekshree/iitk-coin/middleware"
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

	user := db.User{
		Rollno:   p.Rollno,
		Name:     p.Username,
		Password: util.HashAndSalt(p.Password),
	}
	status := db.Add(user)

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

	pass, status := db.FindPass(p.Rollno)
	var res bool = false
	if status {
		res = util.ComparePasswords(pass, p.Password)
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

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t, "status": res})
		}
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": res})
}

func Secret(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	roll := claims["roll"].(string)
	loggedin := claims["loggedin"].(bool)

	// log.Println("The roll is ", roll)

	if loggedin && db.UserExists(roll) {
		return c.Status(fiber.StatusOK).SendString("This is a very secret string.")
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Bad Token, user does not exist.")
}
