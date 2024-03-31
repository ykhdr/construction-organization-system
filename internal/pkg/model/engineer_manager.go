package model

type EngineerManager struct {
	WorkerId int `db:"manager_id"`
	TeamID   int `db:"team_id"`
}
