package controllers

import (
	"casic/src/database"

	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {
	users := database.SelectDBAmbassador(true)
	return c.JSON(users)
}
