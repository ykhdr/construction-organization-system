package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ConstructionTeamRepository interface {
	Save(ctx context.Context, entity model.ConstructionTeam) (int, error)
	Find(ctx context.Context, id int) (*model.ConstructionTeam, error)
	Update(ctx context.Context, entity model.ConstructionTeam) error
	Delete(ctx context.Context, id int) error
}
