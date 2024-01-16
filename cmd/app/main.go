package main

import (
	"log"

	"github.com/ZachHunn/tron/training_platform/internal/api/routes"
	"github.com/ZachHunn/tron/training_platform/internal/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))

	database.DB.Close()
}
