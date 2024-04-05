package model

type EngineerManager struct {
	WorkerId int `db:"manager_id" json:"worker_id"`
	TeamID   int `db:"team_id" json:"team_id"`
}
