package main

import (
	"casic/src/database"
	"casic/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	/*if err := database.AutoMigrate(); err != nil {
		fmt.Println(err)
	}*/
	database.SetupRedis()
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(":3000")
}
