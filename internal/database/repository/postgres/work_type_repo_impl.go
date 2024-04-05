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
