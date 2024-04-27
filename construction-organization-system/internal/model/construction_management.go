package model

type ConstructionManagement struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	ManagerID int    `db:"manager_id" json:"manager_id"`
}
