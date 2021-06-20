package routes

import (
	"github.com/abhishekshree/iitk-coin/config"
	"github.com/abhishekshree/iitk-coin/db"
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

func AwardCoins(c *fiber.Ctx) error {
	user := struct {
		Roll string `json:"rollno"`
		Amt  int    `json:"amount"`
	}{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	res := db.AddCoins(user.Roll, user.Amt)

	if !res {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Failed to award coins.",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Coins Awarded.",
		},
	)
}

func TransferCoins(c *fiber.Ctx) error {
	body := struct {
		From string `json:"from"`
		To   string `json:"to"`
		Amt  int    `json:"amount"`
	}{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	coinsToTransfer := body.Amt - (config.TAX_PERCENT*body.Amt)/100
	res := db.TransferCoins(body.From, body.To, coinsToTransfer)

	if !res {
		return c.Status(fiber.StatusFailedDependency).JSON(
			fiber.Map{
				"message": "Failed to transfer coins.",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Coins Transferred.",
			"amount":  coinsToTransfer,
		},
	)
}
