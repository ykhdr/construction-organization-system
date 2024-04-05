package model

type ConstructionProject struct {
	ID             int    `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	BuildingSiteID int    `db:"building_site_id" json:"building_site_id"`
}
