package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/travboz/fiber-mongo-api/internal/db"
	"github.com/travboz/fiber-mongo-api/internal/routes"
	"github.com/travboz/fiber-mongo-api/pkg/configs"
)

func main() {
	if err := configs.LoadEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db.NewMongoDBInstance("golang-api", "mongodb://localhost:27017")

	app := &application{
		dbInstance: db,
		fiber:      fiber.New(),
	}

	routes.UserRoute(app.fiber)

	app.fiber.Listen(":6000")
}
