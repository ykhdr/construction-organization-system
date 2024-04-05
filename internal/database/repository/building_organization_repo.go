package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type BuildingOrganizationRepository interface {
	Save(ctx context.Context, entity model.BuildingOrganization) (int, error)
	Find(ctx context.Context, id int) (*model.BuildingOrganization, error)
	Update(ctx context.Context, entity model.BuildingOrganization) error
	Delete(ctx context.Context, id int) error
}
