package routes

import (
	"github.com/abhishekshree/iitk-coin/config"
	"github.com/abhishekshree/iitk-coin/db"
	"github.com/abhishekshree/iitk-coin/middleware"
	"github.com/gofiber/fiber/v2"
)

func GetCoins(c *fiber.Ctx) error {
	roll := struct {
		Roll string `json:"rollno"`
	}{}
	if err := c.BodyParser(&roll); err != nil {
		return err
	}

	coins := db.CoinCount(roll.Roll)

	if coins == -1 {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Could not retrieve coins!",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"rollno": roll.Roll,
			"coins":  coins,
		},
	)
}

func isOverflow(coins float64, amount float64) bool {
	return coins+amount > config.MAX_BALANCE
}

func AwardCoins(c *fiber.Ctx) error {
	user := struct {
		Roll string  `json:"rollno"`
		Amt  float64 `json:"amount"`
	}{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	sender_roll := middleware.GetRollFromJWT(c)
	if !db.IsAdmin(sender_roll) {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"message": "Cannot award coins if you're not an admin.",
			},
		)
	}
	if isOverflow(db.CoinCount(user.Roll), user.Amt) {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Cannot award coins, balance overflows.",
			},
		)
	}
	res := db.AddCoins(user.Roll, user.Amt)

	if !res {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Failed to award coins.",
			},
		)
	}

	// Add to transactions table.
	db.AddAwardLog(sender_roll, user.Roll, user.Amt)

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Coins Awarded.",
		},
	)
}

func TransferCoins(c *fiber.Ctx) error {
	body := struct {
		From string  `json:"from"`
		To   string  `json:"to"`
		Amt  float64 `json:"amount"`
	}{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	sender_roll := middleware.GetRollFromJWT(c)
	if sender_roll != body.From {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"message": "User not logged in.",
			},
		)
	}
	if body.From == body.To {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Cannot award oneself.",
			},
		)
	}
	var coinsToTransfer float64

	if body.From[0:2] == body.To[0:2] {
		coinsToTransfer = body.Amt - (config.TAX_PERCENT_INTRABATCH*body.Amt)/100
	} else {
		coinsToTransfer = body.Amt - (config.TAX_PERCENT_INTERBATCH*body.Amt)/100
	}

	if db.CoinCount(body.From) < body.Amt {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Failed to transfer coins, insufficient balance.",
			},
		)
	}

	if isOverflow(db.CoinCount(body.To), body.Amt) {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Cannot transfer coins, balance overflows.",
			},
		)
	}

	res := db.TransferCoins(body.From, body.To, coinsToTransfer)

	if !res {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Failed to transfer coins.",
			},
		)
	}

	// Transfer log
	db.AddTransferLog(body.From, body.To, body.Amt, body.Amt-coinsToTransfer)

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Coins Transferred.",
			"amount":  coinsToTransfer,
		},
	)
}
