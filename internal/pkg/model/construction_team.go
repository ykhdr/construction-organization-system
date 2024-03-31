package model

type ConstructionTeam struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	ProjectID int    `db:"project_id"`
}
