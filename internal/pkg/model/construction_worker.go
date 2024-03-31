package model

type ConstructionWorker struct {
	ID                     int    `db:"id"`
	IsShiftWorker          bool   `db:"is_shift_worker"`
	TeamID                 int    `db:"team_id"`
	Name                   string `db:"name"`
	Surname                string `db:"surname"`
	Patronymic             string `db:"patronymic"`
	Age                    int    `db:"age"`
	Seniority              int    `db:"seniority"`
	BuildingOrganizationID int    `db:"building_organization_id"`
}
