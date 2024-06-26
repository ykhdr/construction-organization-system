package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionManagementRepository interface {
	Save(ctx context.Context, entity model.ConstructionManagement) (int, error)
	Find(ctx context.Context, id int) (*model.ConstructionManagement, error)
	Update(ctx context.Context, entity model.ConstructionManagement) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*model.ConstructionManagement, error)
}
