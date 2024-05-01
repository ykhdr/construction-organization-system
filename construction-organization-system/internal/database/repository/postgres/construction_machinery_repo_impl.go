package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
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

func (repo *constructionMachineryRepository) GetByManagement(ctx context.Context, managementId int) ([]*model.ConstructionMachinery, error) {
	query := `
	SELECT m.id, m.name, m.project_id
	FROM construction_management AS cm
         JOIN building_site AS bs ON cm.id = bs.management_id
         JOIN construction_project AS cp ON bs.id = cp.building_site_id
         JOIN construction_machinery AS m ON cp.id = m.project_id
	WHERE cm.name = $1
  		AND cm.id != 0
	`

	rows, err := repo.db.QueryContext(ctx, query, managementId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.ConstructionMachinery
	for rows.Next() {
		var entity model.ConstructionMachinery
		err = rows.Scan(&entity.ID, &entity.Name, &entity.ProjectID)
		if err != nil {
			return nil, err
		}
		result = append(result, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *constructionMachineryRepository) GetByProject(ctx context.Context, projectID int) ([]*model.ConstructionMachinery, error) {
	query := `
	SELECT cm.id, cm.name, cm.project_id
	FROM construction_project AS cp 
		JOIN construction_machinery AS cm ON cp.id = cm.project_id
	WHERE cp.id = $1
		AND cp.id != 0
	`

	rows, err := repo.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*model.ConstructionMachinery
	for rows.Next() {
		var entity model.ConstructionMachinery
		err = rows.Scan(&entity.ID, &entity.Name, &entity.ProjectID)
		if err != nil {
			return nil, err
		}
		result = append(result, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *constructionMachineryRepository) GetByProjectWithPeriod(ctx context.Context, projectID int, startDate, endDate string) ([]*model.ConstructionMachinery, error) {
	query := `
	SELECT cm.id, cm.name, cm.project_id
	FROM construction_project AS cp 
		JOIN work_schedule AS ws ON cp.id = ws.project_id
        JOIN construction_machinery_usage AS cmu ON ws.id = cmu.work_schedule_id
        JOIN construction_machinery AS cm ON cp.id = cm.project_id
	WHERE cp.id = $1
		AND cp.id != 0
		AND ($2 = '' OR ws.fact_start_date >= to_date($2,'YYYY-MM-DD'))
		AND ($3 = '' OR ws.fact_end_date <= to_date($3, 'YYYY-MM-DD'))
	`

	rows, err := repo.db.QueryContext(ctx, query, projectID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*model.ConstructionMachinery
	for rows.Next() {
		var entity model.ConstructionMachinery
		err = rows.Scan(&entity.ID, &entity.Name, &entity.ProjectID)
		if err != nil {
			return nil, err
		}
		result = append(result, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
