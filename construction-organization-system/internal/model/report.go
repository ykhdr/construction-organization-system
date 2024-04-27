package model

import (
	"encoding/json"
	"time"
)

type Report struct {
	ID                 int             `db:"id" json:"id"`
	ProjectID          int             `db:"project_id" json:"project_id"`
	ReportCreationDate time.Time       `db:"report_creation_date" json:"report_creation_date"`
	ReportFile         json.RawMessage `db:"report_file" json:"report_file"`
}
