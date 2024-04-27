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
	var estimate model.Estimate
	err := repo.db.QueryRowContext(ctx, `
	SELECT e.id, e.last_update_date
	FROM construction_project AS cp
         JOIN estimate AS e ON cp.id = e.id
	WHERE cp.id = $1
  		AND cp.id != 0
	`, id).Scan(&estimate.ID, &estimate.LastUpdateDate)
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.QueryContext(ctx, `
	SELECT estimate_id, material_id, plan_quantity, fact_quantity, name, cost
	FROM material_usage AS mu
		JOIN material AS m ON mu.material_id = m.id
	WHERE estimate_id = $1
		AND estimate_id != 1
	`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var materialsUsages []*model.MaterialUsage
	for rows.Next() {
		var entity model.MaterialUsage
		err := rows.Scan(&entity.EstimateID, &entity.Material.ID, &entity.PlanQuantity, &entity.FactQuantity, &entity.Material.Name, &entity.Material.Cost)
		if err != nil {
			return nil, err
		}

		materialsUsages = append(materialsUsages, &entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	estimate.MaterialUsage = materialsUsages

	return &estimate, nil
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
