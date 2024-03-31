package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type constructionManagerRepository struct {
	db *sql.DB
}

func NewConstructionManagerRepository(db *sql.DB) repository.ConstructionManagerRepository {
	return &constructionManagerRepository{db: db}
}

func (repo *constructionManagerRepository) Save(ctx context.Context, entity model.ConstructionManager) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_manager(worker_id, team_id) VALUES ($1, $2)",
		entity.WorkerID, entity.TeamID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionManagerRepository) FindByTeamId(ctx context.Context, teamId int) (*model.ConstructionManager, error) {
	var entity model.ConstructionManager
	err := repo.db.QueryRowContext(ctx, "SELECT worker_id, team_id FROM construction_manager cm WHERE  team_id = $1", teamId).
		Scan(&entity.WorkerID, &entity.TeamID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionManagerRepository) UpdateByTeamId(ctx context.Context, entity model.ConstructionManager) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_manager SET worker_id = $1 WHERE team_id = $2",
		entity.WorkerID, entity.TeamID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionManagerRepository) Delete(ctx context.Context, workerId, teamId int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_manager WHERE worker_id = $1 AND team_id = $2",
		workerId, teamId)
	if err != nil {
		return err
	}
	return nil
}
