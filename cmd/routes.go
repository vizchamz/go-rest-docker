package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/vizchamz/go-rest-docker/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Default)

    app.Get("/fact", handlers.ListFacts)

    app.Post("/fact", handlers.CreateFact)
}