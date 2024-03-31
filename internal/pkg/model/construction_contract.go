package model

import "time"

type ConstructionContract struct {
	ID                     int       `db:"id"`
	BuildingOrganizationID int       `db:"building_organization_id"`
	CustomerOrganizationID int       `db:"customer_organization_id"`
	Name                   string    `db:"name"`
	ProjectID              int       `db:"project_id"`
	SigningDate            time.Time `db:"signing_date"`
}
