package controllers

import (
	"casic/src/database"
	"casic/src/middlewares"
	"casic/src/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if data["password"] != data["password_confirm"] {
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}

	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: strings.Contains(c.Path(), "/api/ambassador"),
	}
	user.SetPswd(data["password"])
	database.InsertDB(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	user := database.SelectDB(data["email"])
	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if err := user.CheckPswd(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Wrong password",
		})
	}

	IsAmbassador := strings.Contains(c.Path(), "/api/ambassador")

	var scope string

	if IsAmbassador {
		scope = "ambassador"
	} else {
		scope = "admin"
	}

	if (!IsAmbassador && user.IsAmbassador) || (IsAmbassador && !user.IsAmbassador) {

		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	token, err := middlewares.GenerateJWC(user.Id, scope)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Error with jwt",
		})
	}
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func User(c *fiber.Ctx) error {
	id, err := middlewares.GetUserId(c)
	if err != nil {
		return err
	}
	user := database.SelectDBId(id)

	if strings.Contains(c.Path(), "/api/ambassador") {
		ambassador := models.Ambassador(user)
		order := database.SelectDBOrder(user.Id)
		ambassador.CalculateRevenue(order)
		return c.JSON(ambassador)
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	id, err := middlewares.GetUserId(c)

	if err != nil {
		return err
	}
	user := models.User{
		Id:        id,
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	if err = database.UpdateDB(&user); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func UpdatePswd(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	id, err := middlewares.GetUserId(c)
	if err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}
	user := models.User{}
	user.Id = id
	user.SetPswd(data["password"])

	if err = database.UpdateDBPswd(&user); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
