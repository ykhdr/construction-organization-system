package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type ApartmentHouseRepository interface {
	Save(ctx context.Context, entity model.ApartmentHouse) (int, error)
	Find(ctx context.Context, id int) (*model.ApartmentHouse, error)
	Update(ctx context.Context, entity model.ApartmentHouse) error
	Delete(ctx context.Context, id int) error
}
