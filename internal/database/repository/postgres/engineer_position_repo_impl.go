package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type engineerPositionRepository struct {
	db *sql.DB
}

func NewEngineerPositionRepository(db *sql.DB) repository.EngineerPositionRepository {
	return &engineerPositionRepository{db: db}
}

func (repo *engineerPositionRepository) Save(ctx context.Context, entity model.EngineerPosition) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO engineer_position(name) VALUES ($1)",
		entity.Name).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *engineerPositionRepository) Find(ctx context.Context, id int) (*model.EngineerPosition, error) {
	var entity model.EngineerPosition
	err := repo.db.QueryRowContext(ctx, "SELECT id, name FROM engineer_position WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *engineerPositionRepository) Update(ctx context.Context, entity model.EngineerPosition) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE engineer_position SET name = $1 WHERE id = $2",
		entity.Name, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *engineerPositionRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM engineer_position WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
