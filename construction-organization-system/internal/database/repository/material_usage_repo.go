package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type MaterialUsageRepository interface {
	Save(ctx context.Context, entity model.MaterialUsage) (int, error)
	Find(ctx context.Context, id int) (*model.MaterialUsage, error)
	Update(ctx context.Context, entity model.MaterialUsage) error
	Delete(ctx context.Context, id int) error
}
