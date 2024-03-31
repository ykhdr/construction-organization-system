package model

type MaterialUsage struct {
	EstimateID   int `db:"estimate_id"`
	MaterialID   int `db:"material_id"`
	PlanQuantity int `db:"plan_quantity"`
	FactQuantity int `db:"fact_quantity"`
}
