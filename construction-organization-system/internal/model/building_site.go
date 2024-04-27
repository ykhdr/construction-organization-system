package model

type BuildingSite struct {
	ID           int    `db:"id" json:"id"`
	Address      string `db:"address" json:"address"`
	ManagementID int    `db:"management_id" json:"management_id"`
	ManagerID    int    `db:"manager_id" json:"manager_id"`
}
