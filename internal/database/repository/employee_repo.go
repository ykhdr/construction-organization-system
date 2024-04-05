package repository

import (
	"construction-organization-system/internal/model"
	"context"
)

type EmployeeRepository interface {
	Save(ctx context.Context, entity model.Employee) (int, error)
	Find(ctx context.Context, id int) (*model.Employee, error)
	Update(ctx context.Context, entity model.Employee) error
	Delete(ctx context.Context, id int) error
}
