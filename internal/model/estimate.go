package model

import "time"

type Estimate struct {
	ID             int           `db:"id" json:"id"`
	MaterialUsage  MaterialUsage `json:"material_usage"`
	LastUpdateDate time.Time     `db:"last_update_date" json:"last_update_date"`
}
