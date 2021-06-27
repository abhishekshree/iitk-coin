package middleware

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetRollFromJWT(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["roll"].(string)
}
