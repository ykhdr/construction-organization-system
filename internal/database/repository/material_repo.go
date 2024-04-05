package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type MaterialRepository interface {
	Save(ctx context.Context, entity model.Material) (int, error)
	Find(ctx context.Context, id int) (*model.Material, error)
	Update(ctx context.Context, entity model.Material) error
	Delete(ctx context.Context, id int) error
}
