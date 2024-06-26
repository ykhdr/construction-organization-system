package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type ReportRepository interface {
	Save(ctx context.Context, entity model.Report) (int, error)
	Find(ctx context.Context, id int) (*model.Report, error)
	Update(ctx context.Context, entity model.Report) error
	Delete(ctx context.Context, id int) error
	FindByProjectID(ctx context.Context, projectId int) ([]*model.Report, error)
}
