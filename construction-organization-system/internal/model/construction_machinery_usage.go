package model

type ConstructionMachineryUsage struct {
	MachineryID    int `db:"machinery_id" json:"machinery_id"`
	WorkScheduleID int `db:"work_schedule_id" json:"work_schedule_id"`
	Quantity       int `db:"quantity" json:"quantity"`
}
