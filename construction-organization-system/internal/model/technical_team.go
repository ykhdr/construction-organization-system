package model

type TechnicalTeam struct {
	ID                  int    `db:"id" json:"id"`
	IsMaintainMachinery bool   `db:"is_maintain_machinery" json:"is_maintain_machinery"`
	Name                string `db:"name" json:"name"`
	ProjectID           int    `db:"project_id" json:"project_id"`
}
