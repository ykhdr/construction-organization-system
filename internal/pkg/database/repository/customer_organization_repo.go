package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type CustomerOrganizationRepository interface {
	Save(ctx context.Context, entity model.CustomerOrganization) (int, error)
	Find(ctx context.Context, id int) (*model.CustomerOrganization, error)
	Update(ctx context.Context, entity model.CustomerOrganization) error
	Delete(ctx context.Context, id int) error
}
