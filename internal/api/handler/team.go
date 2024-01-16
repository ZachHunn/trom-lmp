package handler

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/ZachHunn/tron/training_platform/internal/database"
	"github.com/ZachHunn/tron/training_platform/models"
	"github.com/gofiber/fiber/v2"
)

type Teams struct {
	Teams []models.Team `json:"teams"`
}

func GetTeams(c *fiber.Ctx) error {
	db := database.DB

	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	query := "SELECT * FROM teams ORDER BY team_id"
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	defer rows.Close()

	teams := Teams{}

	for rows.Next() {
		team := models.Team{}
		if err := rows.Scan(&team.ID, &team.Name, &team.OrgID, &team.Description); err != nil {
			return err
		}

		teams.Teams = append(teams.Teams, team)
	}

	return c.JSON(fiber.Map{"stastus": "success", "message": "Found All teams", "data": teams})
}

func GetTeamByID(c *fiber.Ctx) error {
	db := database.DB
	teamId := c.Params("id")

	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	teamIdInt, err := strconv.Atoi(teamId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invaild Team Id")
	}

	query := "SELECT * FROM teams WHERE team_id = $1"
	row, err := db.Query(query, teamIdInt)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	defer row.Close()

	if row.Next() {
		team := models.Team{}

		if err := row.Scan(&team.ID, &team.Name, &team.OrgID, &team.Description); err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).SendString("Organization Not Found")
			}
			log.Fatal(err)
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Found Team", "data": team})
	}

	return c.Status(fiber.StatusNotFound).SendString("Org Not Found")
}
