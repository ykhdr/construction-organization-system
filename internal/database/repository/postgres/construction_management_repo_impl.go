package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type constructionManagementRepository struct {
	db *sql.DB
}

func NewConstructionManagementRepository(db *sql.DB) repository.ConstructionManagementRepository {
	return &constructionManagementRepository{db: db}
}

func (repo *constructionManagementRepository) Save(ctx context.Context, entity model.ConstructionManagement) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_management(name, manager_id) VALUES ($1, $2)",
		entity.Name, entity.ManagerID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionManagementRepository) Find(ctx context.Context, id int) (*model.ConstructionManagement, error) {
	var entity model.ConstructionManagement
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, manager_id FROM construction_management cm WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.ManagerID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionManagementRepository) Update(ctx context.Context, entity model.ConstructionManagement) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_management SET name = $1, manager_id = $2 WHERE id = $3",
		entity.Name, entity.ManagerID, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionManagementRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_management WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
