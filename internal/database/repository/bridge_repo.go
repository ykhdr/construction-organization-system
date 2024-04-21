package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type BridgeRepository interface {
	Save(ctx context.Context, entity model.Bridge) (int, error)
	Find(ctx context.Context, id int) (*model.Bridge, error)
	Update(ctx context.Context, entity model.Bridge) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*model.Bridge, error)
}
