package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type workTypeRepository struct {
	db *sql.DB
}

func NewWorkTypeRepository(db *sql.DB) repository.WorkTypeRepository {
	return &workTypeRepository{db: db}
}

func (repo *workTypeRepository) Save(ctx context.Context, entity model.WorkType) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO work_type(name) VALUES ($1)",
		entity.Name).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *workTypeRepository) Find(ctx context.Context, id int) (*model.WorkType, error) {
	var entity model.WorkType
	err := repo.db.QueryRowContext(ctx, "SELECT id, name FROM work_type WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *workTypeRepository) Update(ctx context.Context, entity model.WorkType) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE work_type SET name = $1 WHERE id = $2",
		entity.Name, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *workTypeRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM work_type WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *workTypeRepository) FindByProjectWithExceededWorkDeadlines(ctx context.Context, projectID int) ([]*model.WorkType, error) {
	var entities []*model.WorkType

	rows, err := repo.db.QueryContext(ctx, `
	SELECT wt.id, wt.name
	FROM work_schedule AS ws
		JOIN work_type AS wt ON ws.work_type_id = wt.id
	WHERE ws.id != 0
		AND wt.id != 0 
		AND ws.project_id = $1
		AND ws.fact_end_date > ws.plan_end_date
	`, projectID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var entity model.WorkType
		if err := rows.Scan(&entity.ID, &entity.Name); err != nil {
			return nil, err
		}
		entities = append(entities, &entity)
	}

	return entities, nil
}
