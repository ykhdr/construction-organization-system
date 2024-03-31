package model

type ConstructionManager struct {
	WorkerID int `db:"worker_id"`
	TeamID   int `db:"team_id"`
}
