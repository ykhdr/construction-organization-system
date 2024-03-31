package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type constructionMachineryRepository struct {
	db *sql.DB
}

func NewConstructionMachineryRepository(db *sql.DB) repository.ConstructionMachineryRepository {
	return &constructionMachineryRepository{db: db}
}

func (repo *constructionMachineryRepository) Save(ctx context.Context, entity model.ConstructionMachinery) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_machinery(name, project_id) VALUES ($1,$2)",
		entity.Name, entity.ProjectID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionMachineryRepository) Find(ctx context.Context, id int) (*model.ConstructionMachinery, error) {
	var entity model.ConstructionMachinery
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, project_id FROM construction_machinery ah WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.ProjectID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionMachineryRepository) Update(ctx context.Context, entity model.ConstructionMachinery) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_machinery SET name = $1, project_id = $2 WHERE id = $3",
		entity.Name, entity.ProjectID, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionMachineryRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_machinery WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
