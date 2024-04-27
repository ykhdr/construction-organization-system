package model

import "time"

type WorkSchedule struct {
	ID                 int       `db:"id" json:"id"`
	ConstructionTeamID int       `db:"construction_team_id" json:"construction_team_id"`
	WorkType           WorkType  `json:"work_type_id"`
	PlanStartDate      time.Time `db:"plan_start_date" json:"plan_start_date"`
	PlanEndDate        time.Time `db:"plan_end_date" json:"plan_end_date"`
	FactStartDate      time.Time `db:"fact_start_date" json:"fact_start_date"`
	FactEndDate        time.Time `db:"fact_end_date" json:"fact_end_date"`
	PlanOrder          int       `db:"plan_order" json:"plan_order"`
	FactOrder          int       `db:"fact_order" json:"fact_order"`
	ProjectID          int       `db:"project_id" json:"project_id"`
}
