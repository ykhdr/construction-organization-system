package model

type School struct {
	ID             int    `db:"id" json:"id"`
	ClassroomCount int    `db:"classroom_count" json:"classroom_count"`
	Floors         int    `db:"floors" json:"floors"`
	Name           string `db:"name" json:"name"`
	BuildingSiteID int    `db:"building_site_id" json:"building_site_id"`
}
