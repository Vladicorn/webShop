package controllers

import (
	"casic/src/database"

	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	products := database.SelectDBOrders()
	return c.JSON(products)
}
