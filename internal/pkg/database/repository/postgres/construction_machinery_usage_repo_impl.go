package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type constructionMachineryUsageRepository struct {
	db *sql.DB
}

func NewConstructionMachineryUsageRepository(db *sql.DB) repository.ConstructionMachineryUsageRepository {
	return &constructionMachineryUsageRepository{db: db}
}

func (repo *constructionMachineryUsageRepository) Save(ctx context.Context, entity model.ConstructionMachineryUsage) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_machinery_usage(work_schedule_id, machinery_id, quantity) VALUES ($1,$2, $3)",
		entity.WorkScheduleID, entity.MachineryID, entity.Quantity).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionMachineryUsageRepository) Find(ctx context.Context, workScheduleId, machineryId int) (*model.ConstructionMachineryUsage, error) {
	var entity model.ConstructionMachineryUsage
	err := repo.db.QueryRowContext(ctx, "SELECT work_schedule_id, machinery_id, quantity FROM construction_machinery_usage ah WHERE work_schedule_id = $1 AND machinery_id != 0",
		workScheduleId, machineryId).
		Scan(&entity.WorkScheduleID, &entity.MachineryID, &entity.Quantity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionMachineryUsageRepository) Update(ctx context.Context, entity model.ConstructionMachineryUsage) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_machinery_usage SET quantity = $1 WHERE  work_schedule_id = $2 AND machinery_id = $3",
		entity.Quantity, entity.WorkScheduleID, entity.MachineryID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionMachineryUsageRepository) Delete(ctx context.Context, workScheduleId, machineryId int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_machinery_usage WHERE work_schedule_id = $1 AND machinery_id = $2",
		workScheduleId, machineryId)
	if err != nil {
		return err
	}
	return nil
}
