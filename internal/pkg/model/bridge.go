package model

type Bridge struct {
	ID                 int    `db:"id"`
	Span               int    `db:"span"`
	Width              int    `db:"width"`
	TrafficLanesNumber int    `db:"traffic_lanes_number"`
	Name               string `db:"name"`
	BuildingSiteID     int    `db:"building_site_id"`
}
