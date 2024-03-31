package model

import "time"

type Estimate struct {
	ID             int       `db:"id"`
	LastUpdateDate time.Time `db:"last_update_date"`
}
