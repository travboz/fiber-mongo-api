package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/travboz/fiber-mongo-api/internal/db"
	"github.com/travboz/fiber-mongo-api/internal/store"
	"github.com/travboz/fiber-mongo-api/pkg/configs"
)

func main() {
	if err := configs.LoadEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.NewMongoDBInstance("golang-api", "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Error connecting to database.")
	}

	store := store.NewMongoStorage(db)

	app := &application{
		store:     store,
		fiber:     fiber.New(),
		validator: validator.New(),
	}

	app.UserRoutes()

	if err := app.fiber.Listen(":6000"); err != nil {
		log.Fatal(err)
	}
}
