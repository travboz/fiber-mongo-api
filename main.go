package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/travboz/fiber-mongo-api/routes"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	app := fiber.New()

	// app.Get("/", hello)

	// client, err := db.ConnectMongoDB(env.GetString("MONGO_URI", "mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db := client.Database(env.GetString("DBNAME", "golang-api"))

	routes.UserRoute(app)

	app.Listen(":6000")
}
