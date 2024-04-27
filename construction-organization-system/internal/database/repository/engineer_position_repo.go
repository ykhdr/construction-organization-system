package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type EngineerPositionRepository interface {
	Save(ctx context.Context, entity model.EngineerPosition) (int, error)
	Find(ctx context.Context, id int) (*model.EngineerPosition, error)
	Update(ctx context.Context, entity model.EngineerPosition) error
	Delete(ctx context.Context, id int) error
}
