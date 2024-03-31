package model

type BuildingSite struct {
	ID           int    `db:"id"`
	Address      string `db:"address"`
	ManagementID int    `db:"management_id"`
	ManagerID    int    `db:"manager_id"`
}
