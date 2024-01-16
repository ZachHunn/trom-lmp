package models

type Organization struct {
	ID        int     `json:"org_id"`
	Name      string  `json:"org_name"`
	ParentOrg *int    `json:"parentorg"`
	ChildOrg  *int    `json:"childorg"`
	ShortName *string `json:"short_name"`
}

type Team struct {
	ID          int     `json:"team_id"`
	Name        string  `json:"name"`
	OrgID       int     `json:"org_id"`
	Description *string `json:"description"`
}
