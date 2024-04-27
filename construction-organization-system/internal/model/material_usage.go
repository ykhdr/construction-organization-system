package model

type MaterialUsage struct {
	EstimateID   int      `db:"estimate_id" json:"estimate_id"`
	Material     Material `json:"material"`
	PlanQuantity int      `db:"plan_quantity" json:"plan_quantity"`
	FactQuantity int      `db:"fact_quantity" json:"fact_quantity"`
}
