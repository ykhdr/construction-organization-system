package model

type ConstructionWorker struct {
	ID                     int    `db:"id" json:"id"`
	IsShiftWorker          bool   `db:"is_shift_worker" json:"is_shift_worker"`
	TeamID                 int    `db:"team_id" json:"team_id"`
	Name                   string `db:"name" json:"name"`
	Surname                string `db:"surname" json:"surname"`
	Patronymic             string `db:"patronymic" json:"patronymic"`
	Age                    int    `db:"age" json:"age"`
	Seniority              int    `db:"seniority" json:"seniority"`
	BuildingOrganizationID int    `db:"building_organization_id" json:"building_organization_id"`
}
