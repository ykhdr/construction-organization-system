package model

type EngineerWorker struct {
	ID                     int              `db:"id" json:"id"`
	SkillLevel             string           `db:"skill_level" json:"skill_level"`
	Position               EngineerPosition `json:"position"`
	TeamID                 int              `db:"team_id" json:"team_id"`
	Name                   string           `db:"name" json:"name"`
	Surname                string           `db:"surname" json:"surname"`
	Patronymic             string           `db:"patronymic" json:"patronymic"`
	Age                    int              `db:"age" json:"age"`
	Seniority              int              `db:"seniority" json:"seniority"`
	BuildingOrganizationID int              `db:"building_organization_id" json:"building_organization_id"`
}
