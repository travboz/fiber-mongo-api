package main

func (app *application) UserRoutes() {
	app.fiber.Post("/user", app.HandlerCreateUser)
	app.fiber.Get("/user/:userId", app.HandlerGetAUser)
	app.fiber.Put("/user/:userId", app.HandlerEditAUser)
	app.fiber.Delete("/user/:userId", app.HandlerDeleteAUser)
	app.fiber.Get("/users", app.HandlerGetAllUsers)

	app.fiber.Get("/health", app.HandlerHelloHealthCheck)
}
