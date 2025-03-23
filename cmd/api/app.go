package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/travboz/fiber-mongo-api/internal/store"
)

type application struct {
	store     store.Storage
	fiber     *fiber.App
	validator *validator.Validate
}
