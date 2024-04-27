package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type BuildingSiteRepository interface {
	Save(ctx context.Context, entity model.BuildingSite) (int, error)
	Find(ctx context.Context, id int) (*model.BuildingSite, error)
	Update(ctx context.Context, entity model.BuildingSite) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*model.BuildingSite, error)
}
