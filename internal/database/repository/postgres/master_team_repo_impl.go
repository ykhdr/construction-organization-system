package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type masterTeamRepository struct {
	db *sql.DB
}

func NewMasterTeamRepository(db *sql.DB) repository.MasterTeamRepository {
	return &masterTeamRepository{db: db}
}

func (repo *masterTeamRepository) Save(ctx context.Context, entity model.MasterTeam) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO master_team(name, project_id, design_experience_years) VALUES ($1, $2, $3)",
		entity.Name, entity.ProjectID, entity.DesignExperienceYears).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *masterTeamRepository) Find(ctx context.Context, id int) (*model.MasterTeam, error) {
	var entity model.MasterTeam
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, project_id, design_experience_years FROM master_team WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name, &entity.ProjectID, &entity.DesignExperienceYears)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *masterTeamRepository) Update(ctx context.Context, entity model.MasterTeam) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE master_team SET name = $1, project_id = $2, design_experience_years = $3 WHERE id = $4",
		entity.Name, entity.ProjectID, entity.DesignExperienceYears, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *masterTeamRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM master_team WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
