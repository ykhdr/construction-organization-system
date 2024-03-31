package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ConstructionManagerRepository interface {
	Save(ctx context.Context, entity model.ConstructionManager) (int, error)
	FindByTeamId(ctx context.Context, teamId int) (*model.ConstructionManager, error)
	UpdateByTeamId(ctx context.Context, entity model.ConstructionManager) error
	Delete(ctx context.Context, workerId, teamId int) error
}
