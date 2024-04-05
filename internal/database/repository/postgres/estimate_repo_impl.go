package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type estimateRepository struct {
	db *sql.DB
}

func NewEstimateRepository(db *sql.DB) repository.EstimateRepository {
	return &estimateRepository{db: db}
}

func (repo *estimateRepository) Save(ctx context.Context, entity model.Estimate) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO estimate(id, last_update_date) VALUES ($1, $2)",
		entity.ID, entity.LastUpdateDate).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *estimateRepository) Find(ctx context.Context, id int) (*model.Estimate, error) {
	var entity model.Estimate
	err := repo.db.QueryRowContext(ctx, "SELECT id, last_update_date FROM estimate WHERE id = $1", id).
		Scan(&entity.ID, &entity.LastUpdateDate)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *estimateRepository) Update(ctx context.Context, entity model.Estimate) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE estimate SET last_update_date = $1 WHERE id = $2",
		entity.LastUpdateDate, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *estimateRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM estimate WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
