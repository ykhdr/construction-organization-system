package model

type MaterialUsage struct {
	EstimateID   int `db:"estimate_id" json:"estimate_id"`
	MaterialID   int `db:"material_id" json:"material_id"`
	PlanQuantity int `db:"plan_quantity" json:"plan_quantity"`
	FactQuantity int `db:"fact_quantity" json:"fact_quantity"`
}
