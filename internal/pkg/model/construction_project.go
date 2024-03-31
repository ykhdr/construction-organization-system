package model

type ConstructionProject struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	BuildingSiteID int    `db:"buiding_site_id"`
}
