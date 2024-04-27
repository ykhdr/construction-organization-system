package model

type ConstructionMachinery struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	ProjectID int    `db:"project_id" json:"project_id"`
}
