package model

type ApartmentHouse struct {
	ID             int    `db:"id"`
	Floors         int    `db:"floors"`
	Name           string `db:"name"`
	BuildingSiteID int    `db:"building_site_id"`
}
