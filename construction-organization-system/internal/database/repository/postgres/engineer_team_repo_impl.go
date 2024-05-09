package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type engineerTeamRepository struct {
	db *sql.DB
}

func NewEngineerTeamRepository(db *sql.DB) repository.EngineerTeamRepository {
	return &engineerTeamRepository{db: db}
}

func (repo *engineerTeamRepository) Save(ctx context.Context, entity model.EngineerTeam) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO engineer_team(name, project_id, type) VALUES ($1, $2)",
		entity.Name, entity.ProjectID, entity.Type).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *engineerTeamRepository) Find(ctx context.Context, id int) (*model.EngineerTeam, error) {
	var entity model.EngineerTeam
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, project_id, type FROM engineer_team WHERE id = $1 and id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.ProjectID, &entity.Type)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *engineerTeamRepository) Update(ctx context.Context, entity model.EngineerTeam) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE engineer_team SET name = $1, project_id = $2 WHERE id = $3",
		entity.Name, entity.ProjectID, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *engineerTeamRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM engineer_team WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *engineerTeamRepository) FindAll(ctx context.Context) ([]*model.EngineerTeam, error) {
	var entities []*model.EngineerTeam
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, project_id, type FROM engineer_team WHERE id != 0")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var entity model.EngineerTeam
		err := rows.Scan(&entity.ID, &entity.Name, &entity.ProjectID, &entity.Type)
		if err != nil {
			return nil, err
		}
		entities = append(entities, &entity)
	}
	return entities, nil
}
