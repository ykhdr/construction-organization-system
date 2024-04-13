package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type constructionTeamRepository struct {
	db *sql.DB
}

func NewConstructionTeamRepository(db *sql.DB) repository.ConstructionTeamRepository {
	return &constructionTeamRepository{db: db}
}

func (repo *constructionTeamRepository) Save(ctx context.Context, entity model.ConstructionTeam) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_team(project_id) VALUES ($1)",
		entity.ProjectID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionTeamRepository) Find(ctx context.Context, id int) (*model.ConstructionTeam, error) {
	var entity model.ConstructionTeam
	err := repo.db.QueryRowContext(ctx, "SELECT id,name, project_id FROM construction_team WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name, &entity.ProjectID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionTeamRepository) Update(ctx context.Context, entity model.ConstructionTeam) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_team SET project_id = $1 WHERE id = $2",
		entity.ProjectID, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionTeamRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_team WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionTeamRepository) FindByProject(ctx context.Context, projectID int) ([]*model.ConstructionTeam, error) {
	query := `
		SELECT ct.id, ct.name, ct.project_id
		FROM construction_project AS cp 
			JOIN work_schedule AS ws ON cp.id = ws.project_id
			JOIN construction_team AS ct ON ct.project_id = cp.id
		WHERE cp.id = $1
		AND ws.fact_start_date <= current_date
	`

	rows, err := repo.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var teams []*model.ConstructionTeam
	for rows.Next() {
		var entity model.ConstructionTeam
		err := rows.Scan(&entity.ID, &entity.Name, &entity.ProjectID)
		if err != nil {
			return nil, err
		}
		teams = append(teams, &entity)
	}
	return teams, nil
}
