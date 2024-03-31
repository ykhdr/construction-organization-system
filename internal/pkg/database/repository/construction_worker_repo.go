package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ConstructionWorkerRepository interface {
	Save(ctx context.Context, entity model.ConstructionWorker) (int, error)
	Find(ctx context.Context, id int) (*model.ConstructionWorker, error)
	Update(ctx context.Context, entity model.ConstructionWorker) error
	Delete(ctx context.Context, id int) error
}
