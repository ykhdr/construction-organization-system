package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type SchoolRepository interface {
	Save(ctx context.Context, entity model.School) (int, error)
	Find(ctx context.Context, id int) (*model.School, error)
	Update(ctx context.Context, entity model.School) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*model.School, error)
}
