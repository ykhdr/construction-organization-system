package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type materialRepository struct {
	db *sql.DB
}

func NewMaterialRepository(db *sql.DB) repository.MaterialRepository {
	return &materialRepository{db: db}
}

func (repo *materialRepository) Save(ctx context.Context, entity model.Material) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO material(name, cost) VALUES ($1, $2)",
		entity.Name, entity.Cost).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *materialRepository) Find(ctx context.Context, id int) (*model.Material, error) {
	var entity model.Material
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, cost FROM material WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.Cost)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *materialRepository) Update(ctx context.Context, entity model.Material) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE material SET name = $1, cost = $2 WHERE id = $3",
		entity.Name, entity.Cost, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *materialRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM material WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *materialRepository) FindByEstimateWithExceededUsage(ctx context.Context, estimateID int) ([]*model.Material, error) {
	var entities []*model.Material

	rows, err := repo.db.QueryContext(ctx, `
	SELECT m.id, m.name, m.cost
	FROM estimate AS e 
		JOIN material_usage AS mu ON mu.estimate_id = e.id
		JOIN material AS m ON mu.material_id = m.id
	WHERE e.id = $1 
		AND mu.fact_quantity > mu.plan_quantity
	`, estimateID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var entity model.Material
		err = rows.Scan(&entity.ID, &entity.Name, &entity.Cost)
		if err != nil {
			return nil, err
		}
		entities = append(entities, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
