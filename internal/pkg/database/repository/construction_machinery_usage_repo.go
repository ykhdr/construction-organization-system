package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ConstructionMachineryUsageRepository interface {
	Save(ctx context.Context, entity model.ConstructionMachineryUsage) (int, error)
	Find(ctx context.Context, workScheduleId, machineryId int) (*model.ConstructionMachineryUsage, error)
	Update(ctx context.Context, entity model.ConstructionMachineryUsage) error
	Delete(ctx context.Context, workScheduleId, machineryId int) error
}
