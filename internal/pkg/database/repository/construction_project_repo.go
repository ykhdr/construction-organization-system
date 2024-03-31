package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ConstructionProjectRepository interface {
	Save(ctx context.Context, entity model.ConstructionProject) (int, error)
	Find(ctx context.Context, id int) (*model.ConstructionProject, error)
	Update(ctx context.Context, entity model.ConstructionProject) error
	Delete(ctx context.Context, id int) error
}
