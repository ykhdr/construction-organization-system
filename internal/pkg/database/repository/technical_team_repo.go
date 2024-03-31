package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type TechnicalTeamRepository interface {
	Save(ctx context.Context, entity model.TechnicalTeam) (int, error)
	Find(ctx context.Context, id int) (*model.TechnicalTeam, error)
	Update(ctx context.Context, entity model.TechnicalTeam) error
	Delete(ctx context.Context, id int) error
}
