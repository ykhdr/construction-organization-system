package model

type BuildingOrganization struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
