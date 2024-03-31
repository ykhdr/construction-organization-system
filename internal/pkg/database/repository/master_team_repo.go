package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type MasterTeamRepository interface {
	Save(ctx context.Context, entity model.MasterTeam) (int, error)
	Find(ctx context.Context, id int) (*model.MasterTeam, error)
	Update(ctx context.Context, entity model.MasterTeam) error
	Delete(ctx context.Context, id int) error
}
