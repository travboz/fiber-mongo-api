package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/travboz/fiber-mongo-api/internal/db"
)

type application struct {
	dbInstance *db.MongoDBInstance
	fiber      *fiber.App
}
