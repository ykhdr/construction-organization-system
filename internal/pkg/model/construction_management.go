package model

type ConstructionManagement struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	ManagerID int    `db:"manager_id"`
}
