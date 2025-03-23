package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/travboz/fiber-mongo-api/controllers"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)

	app.Get("/health", controllers.HelloHealthCheck)
}
