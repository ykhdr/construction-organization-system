package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type WorkScheduleRepository interface {
	Save(ctx context.Context, entity model.WorkSchedule) (int, error)
	Find(ctx context.Context, id int) (*model.WorkSchedule, error)
	Update(ctx context.Context, entity model.WorkSchedule) error
	Delete(ctx context.Context, id int) error
}
