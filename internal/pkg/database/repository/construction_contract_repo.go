package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ConstructionContractRepository interface {
	Save(ctx context.Context, entity model.ConstructionContract) (int, error)
	Find(ctx context.Context, id int) (*model.ConstructionContract, error)
	Update(ctx context.Context, entity model.ConstructionContract) error
	Delete(ctx context.Context, id int) error
}
