package model

type TechnicalTeam struct {
	ID                  int    `db:"id"`
	IsMaintainMachinery bool   `db:"is_maintain_machinery"`
	Name                string `db:"name"`
	ProjectID           int    `db:"project_id"`
}
