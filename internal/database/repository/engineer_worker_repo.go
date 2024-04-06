package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type EngineerWorkerRepository interface {
	Save(ctx context.Context, entity model.EngineerWorker) (int, error)
	Find(ctx context.Context, id int) (*model.EngineerWorker, error)
	Update(ctx context.Context, entity model.EngineerWorker) error
	Delete(ctx context.Context, id int) error
	FindByBuildingSite(ctx context.Context, buildingSiteId int) ([]*model.EngineerWorker, error)
	FindByConstructionManagement(ctx context.Context, constructionManagementId int) ([]*model.EngineerWorker, error)
}
