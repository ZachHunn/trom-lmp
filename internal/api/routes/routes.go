package routes

import (
	"github.com/ZachHunn/tron/training_platform/internal/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	//Organization

	organization := api.Group("/organization")
	organization.Get("/", handler.GetAllOrgs)
	organization.Get("/:id", handler.GetOrgByID)

	//Teams

	team := api.Group("/team")
	team.Get("/", handler.GetTeams)
	team.Get("/:id", handler.GetTeamByID)
}
