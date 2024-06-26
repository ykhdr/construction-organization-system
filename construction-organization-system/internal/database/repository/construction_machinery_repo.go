package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionMachineryRepository interface {
	Save(ctx context.Context, entity model.ConstructionMachinery) (int, error)
	Find(ctx context.Context, id int) (*model.ConstructionMachinery, error)
	Update(ctx context.Context, entity model.ConstructionMachinery) error
	Delete(ctx context.Context, id int) error
	GetByManagement(ctx context.Context, managementId int) ([]*model.ConstructionMachinery, error)
	GetByProject(ctx context.Context, projectID int) ([]*model.ConstructionMachinery, error)
	GetByProjectWithPeriod(ctx context.Context, projectID int, startDate string, endDate string) ([]*model.ConstructionMachinery, error)
	FindAll(ctx context.Context) ([]*model.ConstructionMachinery, error)
}
