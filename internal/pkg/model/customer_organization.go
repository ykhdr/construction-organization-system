package model

type CustomerOrganization struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
