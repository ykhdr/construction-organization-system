package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type EngineerManagerRepository interface {
	Save(ctx context.Context, entity model.EngineerManager) (int, error)
	FindByTeamId(ctx context.Context, teamId int) (*model.EngineerManager, error)
	UpdateByTeamId(ctx context.Context, entity model.EngineerManager) error
	Delete(ctx context.Context, workerId, teamId int) error
}
