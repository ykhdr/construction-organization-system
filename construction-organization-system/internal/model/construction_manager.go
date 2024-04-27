package model

type ConstructionManager struct {
	WorkerID int `db:"worker_id" json:"worker_id"`
	TeamID   int `db:"team_id" json:"team_id"`
}
