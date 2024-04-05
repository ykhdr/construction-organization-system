package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type technicalTeamRepository struct {
	db *sql.DB
}

func NewTechnicalTeamRepository(db *sql.DB) repository.TechnicalTeamRepository {
	return &technicalTeamRepository{db: db}
}

func (repo *technicalTeamRepository) Save(ctx context.Context, entity model.TechnicalTeam) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO technical_team(name, project_id, is_maintain_machinery)  VALUES ($1, $2, $3)",
		entity.Name, entity.ProjectID, entity.IsMaintainMachinery).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *technicalTeamRepository) Find(ctx context.Context, id int) (*model.TechnicalTeam, error) {
	var entity model.TechnicalTeam
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, project_id, is_maintain_machinery FROM technical_team WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name, &entity.ProjectID, &entity.IsMaintainMachinery)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *technicalTeamRepository) Update(ctx context.Context, entity model.TechnicalTeam) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE technical_team SET name = $1, project_id = $2, is_maintain_machinery = $3 WHERE id = $4",
		entity.Name, entity.ProjectID, entity.IsMaintainMachinery, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *technicalTeamRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM technical_team WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
