package routes

import (
	"casic/src/controllers"
	"casic/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuth := admin.Use(middlewares.Auth)
	adminAuth.Get("user", controllers.User)
	adminAuth.Get("logout", controllers.Logout)
	adminAuth.Put("users/info", controllers.UpdateInfo)
	adminAuth.Put("users/update", controllers.UpdatePswd)
	adminAuth.Get("ambassadors", controllers.Ambassadors)

	adminAuth.Get("products", controllers.Products)
	adminAuth.Post("product", controllers.CreateProduct)
	adminAuth.Get("product/:id?", controllers.GetProduct)
	adminAuth.Put("product/:id?", controllers.UpdateProduct)
	adminAuth.Delete("product/:id?", controllers.DeleteProduct)
	adminAuth.Get("orders", controllers.Orders)

	ambassador := api.Group("ambassador")
	ambassador.Post("register", controllers.Register)
	ambassador.Post("login", controllers.Login)
	ambassador.Get("products/frontend", controllers.ProductFont)
	ambassador.Get("products/backend", controllers.ProductBackend)

	ambassadorAuth := ambassador.Use(middlewares.Auth)
	ambassadorAuth.Get("user", controllers.User)
	ambassadorAuth.Get("logout", controllers.Logout)
	ambassadorAuth.Put("users/info", controllers.UpdateInfo)
	ambassadorAuth.Put("users/update", controllers.UpdatePswd)

}
