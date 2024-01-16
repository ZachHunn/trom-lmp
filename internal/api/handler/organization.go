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

type Organizations struct {
	Organizations []models.Organization `json:"organizations"`
}

func GetAllOrgs(c *fiber.Ctx) error {
	db := database.DB

	if db == nil {
		return fmt.Errorf("database connection failed")
	}

	rows, err := db.Query("SELECT * FROM organizations ORDER BY org_id;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	defer rows.Close()
	orgs := Organizations{}

	for rows.Next() {
		org := models.Organization{}
		if err := rows.Scan(&org.ID, &org.Name, &org.ShortName, &org.ParentOrg, &org.ChildOrg); err != nil {
			return err
		}

		orgs.Organizations = append(orgs.Organizations, org)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "All Organizations", "data": orgs})
}

func GetOrgByID(c *fiber.Ctx) error {
	db := database.DB
	orgID := c.Params("id")

	if db == nil {
		return fmt.Errorf("Database connection has failed")
	}

	orgIDInt, err := strconv.Atoi(orgID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invaild Org Id")
	}

	query := "SELECT * FROM organizations WHERE org_id = $1"

	row, err := db.Query(query, orgIDInt)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	defer row.Close()

	if row.Next() {
		org := models.Organization{}

		if err := row.Scan(&org.ID, &org.Name, &org.ParentOrg, &org.ChildOrg, &org.ShortName); err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).SendString("Organization Not Found")
			}
			log.Fatal(err)
		}

		return c.JSON(fiber.Map{"status": "succuess", "message": "Org Found", "data": org})

	}

	return c.Status(fiber.StatusNotFound).SendString("Org Not Found")
}
