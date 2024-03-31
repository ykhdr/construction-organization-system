package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type constructionProjectRepository struct {
	db *sql.DB
}

func NewConstructionProjectRepository(db *sql.DB) repository.ConstructionProjectRepository {
	return &constructionProjectRepository{db: db}
}

func (repo *constructionProjectRepository) Save(ctx context.Context, entity model.ConstructionProject) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_project(name, building_site_id) VALUES ($1, $2); INSERT INTO estimate(id, last_update_date) VALUES (LASTVAL(), NOW()) RETURNING id",
		entity.Name, entity.BuildingSiteID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionProjectRepository) Find(ctx context.Context, id int) (*model.ConstructionProject, error) {
	var entity model.ConstructionProject
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, building_site_id FROM construction_project WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionProjectRepository) Update(ctx context.Context, entity model.ConstructionProject) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_project SET name = $1, building_site_id = $2 WHERE id = $3",
		entity.Name, entity.BuildingSiteID, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionProjectRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_project WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
