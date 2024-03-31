package model

type EngineerWorker struct {
	ID                     int    `db:"id"`
	SkillLevel             string `db:"skill_level"`
	PositionID             int    `db:"position_id"`
	TeamID                 int    `db:"team_id"`
	Name                   string `db:"name"`
	Surname                string `db:"surname"`
	Patronymic             string `db:"patronymic"`
	Age                    int    `db:"age"`
	Seniority              int    `db:"seniority"`
	BuildingOrganizationID int    `db:"building_organization_id"`
}
