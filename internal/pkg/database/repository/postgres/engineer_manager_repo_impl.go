package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type engineerManagerRepository struct {
	db *sql.DB
}

func NewEngineerManagerRepository(db *sql.DB) repository.EngineerManagerRepository {
	return &engineerManagerRepository{db: db}
}

func (repo *engineerManagerRepository) Save(ctx context.Context, entity model.EngineerManager) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO engineer_manager(worker_id, team_id)  VALUES ($1, $2) ",
		entity.WorkerId, entity.TeamID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *engineerManagerRepository) FindByTeamId(ctx context.Context, teamId int) (*model.EngineerManager, error) {
	var entity model.EngineerManager
	err := repo.db.QueryRowContext(ctx, "SELECT worker_id, team_id FROM engineer_manager WHERE team_id = $1 AND team_id != 0", teamId).
		Scan(&entity.WorkerId, &entity.TeamID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *engineerManagerRepository) UpdateByTeamId(ctx context.Context, entity model.EngineerManager) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE engineer_manager SET worker_id = $1 WHERE team_id = $3",
		entity.WorkerId, entity.TeamID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *engineerManagerRepository) Delete(ctx context.Context, workerId, teamId int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM engineer_manager WHERE worker_id = $1 AND team_id = $2",
		workerId, teamId)
	if err != nil {
		return err
	}
	return nil
}
