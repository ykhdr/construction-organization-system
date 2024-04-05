package model

type Employee struct {
	ID                     int    `db:"id" json:"id"`
	Name                   string `db:"name" json:"name"`
	Surname                string `db:"surname" json:"surname"`
	Patronymic             string `db:"patronymic" json:"patronymic"`
	Age                    int    `db:"age" json:"age"`
	Seniority              int    `db:"seniority" json:"seniority"`
	BuildingOrganizationID int    `db:"building_organization_id" json:"building_organization_id"`
}
