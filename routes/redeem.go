package routes

import (
	"github.com/abhishekshree/iitk-coin/config"
	"github.com/abhishekshree/iitk-coin/db"
	"github.com/abhishekshree/iitk-coin/middleware"
	"github.com/gofiber/fiber/v2"
)

func GetRedeemList(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(config.REDEEM_LIST)
}

func RequestItem(c *fiber.Ctx) error {
	req := struct {
		Item string `json:"item"`
	}{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	roll := middleware.GetRollFromJWT(c)
	if db.CoinCount(roll) < config.REDEEM_LIST[req.Item] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Not enough coins",
		},
		)
	} else {
		ok := db.AddRedeemRequest(roll, req.Item)
		if ok {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Success",
			},
			)
		} else {
			return c.Status(fiber.StatusFailedDependency).JSON(fiber.Map{
				"message": "Failed to add a redeem request.",
			},
			)
		}

	}
}

func RejectRedeemRequest(c *fiber.Ctx) error {
	req := struct {
		ID int `json:"id"`
	}{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	roll := middleware.GetRollFromJWT(c)
	if !db.IsAdmin(roll) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Admin only route",
		},
		)
	}
	if db.RejectRedeemRequest(req.ID) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
		},
		)
	} else {
		return c.Status(fiber.StatusFailedDependency).JSON(fiber.Map{
			"message": "Failed to reject a redeem request.",
		},
		)
	}
}

func AcceptRedeemRequest(c *fiber.Ctx) error {
	req := struct {
		ID int `json:"id"`
	}{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	roll := middleware.GetRollFromJWT(c)
	if !db.IsAdmin(roll) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Admin only route",
		},
		)
	}
	rollno, item, status := db.GetRedeemRequest(req.ID)
	if db.CoinCount(rollno) < config.REDEEM_LIST[item] || status != 0 {
		return c.Status(fiber.StatusFailedDependency).JSON(fiber.Map{
			"message": "Failed to accept a redeem request.",
		},
		)
	}
	if db.AcceptRedeemRequest(req.ID) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
		},
		)
	} else {
		return c.Status(fiber.StatusFailedDependency).JSON(fiber.Map{
			"message": "Failed to accept a redeem request.",
		},
		)
	}
}

func RejectPendingRequests(c *fiber.Ctx) error {
	roll := middleware.GetRollFromJWT(c)
	if !db.IsAdmin(roll) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Admin only route",
		},
		)
	}
	req := struct {
		Rollno string `json:"roll"`
	}{}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if db.RejectPendingRedeemRequests(req.Rollno) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
		},
		)
	} else {
		return c.Status(fiber.StatusFailedDependency).JSON(fiber.Map{
			"message": "Failed to reject pending requests.",
		},
		)
	}
}
