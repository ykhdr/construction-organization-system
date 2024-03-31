package model

type Employee struct {
	ID                     int    `db:"id"`
	Name                   string `db:"name"`
	Surname                string `db:"surname"`
	Patronymic             string `db:"patronymic"`
	Age                    int    `db:"age"`
	Seniority              int    `db:"seniority"`
	BuildingOrganizationID int    `db:"building_organization_id"`
}
