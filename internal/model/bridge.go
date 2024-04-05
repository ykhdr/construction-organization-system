package model

type Bridge struct {
	ID                 int    `db:"id" json:"id"`
	Span               int    `db:"span" json:"span"`
	Width              int    `db:"width" json:"width"`
	TrafficLanesNumber int    `db:"traffic_lanes_number" json:"traffic_lanes_number"`
	Name               string `db:"name" json:"name"`
	BuildingSiteID     int    `db:"building_site_id" json:"building_site_id"`
}
