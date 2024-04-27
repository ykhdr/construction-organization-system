package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type EstimateRepository interface {
	Save(ctx context.Context, entity model.Estimate) (int, error)
	Find(ctx context.Context, id int) (*model.Estimate, error)
	Update(ctx context.Context, entity model.Estimate) error
	Delete(ctx context.Context, id int) error
}
