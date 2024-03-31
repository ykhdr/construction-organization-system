package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ConstructionMachineryRepository interface {
	Save(ctx context.Context, entity model.ConstructionMachinery) (int, error)
	Find(ctx context.Context, id int) (*model.ConstructionMachinery, error)
	Update(ctx context.Context, entity model.ConstructionMachinery) error
	Delete(ctx context.Context, id int) error
}
