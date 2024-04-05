package model

import "time"

type ConstructionContract struct {
	ID                     int       `db:"id" json:"id"`
	BuildingOrganizationID int       `db:"building_organization_id" json:"building_organization_id"`
	CustomerOrganizationID int       `db:"customer_organization_id" json:"customer_organization_id"`
	Name                   string    `db:"name" json:"name"`
	ProjectID              int       `db:"project_id" json:"project_id"`
	SigningDate            time.Time `db:"signing_date" json:"signing_date"`
}
