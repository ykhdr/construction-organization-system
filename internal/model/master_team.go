package model

type MasterTeam struct {
	ID                    int    `db:"id" json:"id"`
	DesignExperienceYears int    `db:"design_experience_years" json:"design_experience_years"`
	Name                  string `db:"name" json:"name"`
	ProjectID             int    `db:"project_id" json:"project_id"`
}
