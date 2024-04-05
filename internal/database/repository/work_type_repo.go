package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type WorkTypeRepository interface {
	Save(ctx context.Context, entity model.WorkType) (int, error)
	Find(ctx context.Context, id int) (*model.WorkType, error)
	Update(ctx context.Context, entity model.WorkType) error
	Delete(ctx context.Context, id int) error
}
