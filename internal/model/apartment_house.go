package model

type ApartmentHouse struct {
	ID             int    `db:"id" json:"id"`
	Floors         int    `db:"floors" json:"floors"`
	Name           string `db:"name" json:"name"`
	BuildingSiteID int    `db:"building_site_id" json:"building_site_id"`
}
