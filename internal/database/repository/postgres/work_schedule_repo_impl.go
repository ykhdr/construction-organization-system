package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type workScheduleRepository struct {
	db *sql.DB
}

func NewWorkScheduleRepository(db *sql.DB) repository.WorkScheduleRepository {
	return &workScheduleRepository{db: db}
}

func (repo *workScheduleRepository) Save(ctx context.Context, entity model.WorkSchedule) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO work_schedule(construction_team_id, work_type_id, plan_start_date, plan_end_date, fact_start_date, fact_end_date, plan_order, fact_order, project_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		entity.ConstructionTeamID, entity.WorkTypeID, entity.PlanStartDate, entity.PlanEndDate, entity.FactStartDate, entity.FactEndDate, entity.PlanOrder, entity.FactOrder, entity.ProjectID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *workScheduleRepository) Find(ctx context.Context, id int) (*model.WorkSchedule, error) {
	var entity model.WorkSchedule
	err := repo.db.QueryRowContext(ctx, "SELECT id, construction_team_id, work_type_id, plan_start_date, plan_end_date, fact_start_date, fact_end_date, plan_order, fact_order, project_id FROM work_schedule WHERE id = $1", id).
		Scan(&entity.ID, &entity.ConstructionTeamID, &entity.WorkTypeID, &entity.PlanStartDate, &entity.PlanEndDate, &entity.FactStartDate, &entity.FactEndDate, &entity.PlanOrder, &entity.FactOrder, &entity.ProjectID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *workScheduleRepository) Update(ctx context.Context, entity model.WorkSchedule) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE work_schedule SET construction_team_id = $1, work_type_id = $2, plan_start_date = $3, plan_end_date = $4, fact_start_date = $5, fact_end_date = $6, plan_order = $7, fact_order = $8, project_id = $9 WHERE id = $10",
		entity.ConstructionTeamID, entity.WorkTypeID, entity.PlanStartDate, entity.PlanEndDate, entity.FactStartDate, entity.FactEndDate, entity.PlanOrder, entity.FactOrder, entity.ProjectID, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *workScheduleRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM work_schedule WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
