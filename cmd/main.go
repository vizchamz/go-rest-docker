package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "github.com/vizchamz/go-rest-docker/database"
    "log"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found or could not be loaded")
    }

    database.ConnectDb()
    app := fiber.New()

    setupRoutes(app)

    app.Listen(":3000")
}