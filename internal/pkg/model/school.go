package model

type School struct {
	ID             int    `db:"id"`
	ClassroomCount int    `db:"classroom_count"`
	Floors         int    `db:"floors"`
	Name           string `db:"name"`
	BuildingSiteID int    `db:"building_site_id"`
}
