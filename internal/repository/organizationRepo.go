package repository

import (
	"fmt"

	"github.com/ZachHunn/tron/training_platform/internal/database"
	"github.com/ZachHunn/tron/training_platform/models"
)

type Organizations struct {
	Organizations []models.Organization `json:"organizations"`
}

func GetAllOrganizations() ([]models.Organization, error) {
	db := database.DB

	if db == nil {
		return nil, fmt.Errorf("databasre connection failed")
	}

	rows, err := db.Query("SELECT * FROM organizations order by org_id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := Organizations{}

	for rows.Next() {
		org := models.Organization{}
		if err := rows.Scan(&org.ID, &org.Name, &org.ShortName, &org.ParentOrg, &org.ChildOrg); err != nil {
			return nil, err
		}

		results.Organizations = append(results.Organizations, org)
		fmt.Println(results.Organizations)
	}

	return results.Organizations, nil
}
