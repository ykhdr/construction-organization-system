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
	query := `
	SELECT e.id, m.name, m.cost, mu.plan_quantity, mu.fact_quantity, e.last_update_date
	FROM construction_project AS cp
         JOIN estimate AS e ON cp.id = e.id
         JOIN material_usage AS mu ON e.id = mu.estimate_id
         JOIN material AS m ON mu.material_id = m.id
	WHERE cp.id = $1
  		AND cp.id != 0
	`

	var entity model.Estimate
	err := repo.db.QueryRowContext(ctx, query, id).
		Scan(&entity.ID, &entity.MaterialUsage.Material.Name, &entity.MaterialUsage.Material.Cost, &entity.MaterialUsage.PlanQuantity, &entity.MaterialUsage.FactQuantity, &entity.LastUpdateDate)

	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *estimateRepository) Update(ctx context.Context, entity model.Estimate) error {
	// TODO прописать обновление MaterialUsage в том числе
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
