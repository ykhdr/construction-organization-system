package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
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

func (repo *constructionProjectRepository) FindByConstructionManagement(ctx context.Context, managementID int) ([]*model.ConstructionProject, error) {
	var projects []*model.ConstructionProject

	rows, err := repo.db.QueryContext(ctx, `
	SELECT cp.id, cp.name, cp.building_site_id 
	FROM construction_management AS cm
        JOIN building_site AS bs ON cm.id = bs.management_id
        JOIN construction_project AS cp ON bs.id = cp.building_site_id
	WHERE cm.id = $1
	AND cm.id != 0
		
	`, managementID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.ConstructionProject
		if err := rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID); err != nil {
			return nil, err
		}
		projects = append(projects, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (repo *constructionProjectRepository) FindByBuildingSite(ctx context.Context, siteID int) ([]*model.ConstructionProject, error) {
	var projects []*model.ConstructionProject

	rows, err := repo.db.QueryContext(ctx, `
	SELECT cp.id, cp.name, cp.building_site_id 
    FROM building_site AS bs 
        JOIN construction_project AS cp ON bs.id = cp.building_site_id
	WHERE bs.id = $1
	AND bs.id != 0
		
	`, siteID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.ConstructionProject
		if err := rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID); err != nil {
			return nil, err
		}
		projects = append(projects, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (repo *constructionProjectRepository) FindAll(ctx context.Context) ([]*model.ConstructionProject, error) {
	var projects []*model.ConstructionProject

	rows, err := repo.db.QueryContext(ctx, `
	SELECT id, name, building_site_id
	FROM construction_project
	WHERE id != 0
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.ConstructionProject
		if err := rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID); err != nil {
			return nil, err
		}
		projects = append(projects, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, err
}

func (repo *constructionProjectRepository) FindByWorkTypeWithPeriod(ctx context.Context, workTypeID int, startDate string, endDate string) ([]*model.ConstructionProject, error) {
	var projects []*model.ConstructionProject

	rows, err := repo.db.QueryContext(ctx, `
	SELECT cp.id, cp.name, cp.building_site_id
	FROM construction_project AS cp 
		JOIN work_schedule AS ws ON cp.id = ws.project_id
	WHERE cp.id != 0
		AND ($1 = 0 OR ws.work_type_id = $1)
		AND ($2 = '' OR ws.fact_start_date >= $2)
		AND ($3 = '' OR ws.fact_end_date <= $3)
	`, workTypeID, startDate, endDate)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.ConstructionProject
		if err := rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID); err != nil {
			return nil, err
		}
		projects = append(projects, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, err
}

func (repo *constructionProjectRepository) FindByOrganization(ctx context.Context, organizationID int) ([]*model.ConstructionProject, error) {
	var projects []*model.ConstructionProject

	rows, err := repo.db.QueryContext(ctx, `
	SELECT cp.id, cp.name, cp.building_site_id
	FROM building_organization AS bo
		JOIN construction_contract AS cc ON bo.id = cc.building_organization_id
		JOIN construction_project AS cp ON cc.project_id = cp.id
	WHERE bo.id = $1
	AND bo.id != 0

	`, organizationID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.ConstructionProject
		if err := rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID); err != nil {
			return nil, err
		}
		projects = append(projects, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
