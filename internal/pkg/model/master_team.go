package model

type MasterTeam struct {
	ID                    int    `db:"id"`
	DesignExperienceYears int    `db:"design_experience_years"`
	Name                  string `db:"name"`
	ProjectID             int    `db:"project_id"`
}
