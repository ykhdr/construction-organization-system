package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type engineerWorkerRepository struct {
	db *sql.DB
}

func NewEngineerWorkerRepository(db *sql.DB) repository.EngineerWorkerRepository {
	return &engineerWorkerRepository{db: db}
}

func (repo *engineerWorkerRepository) Save(ctx context.Context, entity model.EngineerWorker) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO engineer_worker(name, surname, patronymic, age, seniority, building_organization_id, team_id, position_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID, entity.TeamID, entity.Position.ID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *engineerWorkerRepository) Find(ctx context.Context, id int) (*model.EngineerWorker, error) {
	var entity model.EngineerWorker

	err := repo.db.QueryRowContext(ctx, `
		SELECT ew.id, ew.name, ew.surname, ew.patronymic, ew.age, ew.seniority, ew.building_organization_id, ew.team_id, ep.id, ep.name 
		FROM engineer_worker AS ew
		JOIN engineer_position AS ep ON ep.id = ew.position_id
		WHERE ew.id = $1
	`, id).
		Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Seniority, &entity.BuildingOrganizationID, &entity.TeamID, &entity.Position.ID, &entity.Position.Name)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *engineerWorkerRepository) Update(ctx context.Context, entity model.EngineerWorker) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE engineer_worker SET name = $1, surname = $2, patronymic = $3, age = $4, seniority = $5, building_organization_id = $6, team_id = $8, position_id = $9 WHERE id = $10",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID, entity.Position.ID, entity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *engineerWorkerRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM engineer_worker WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *engineerWorkerRepository) FindByBuildingSite(ctx context.Context, buildingSiteId int) ([]*model.EngineerWorker, error) {
	var engineerWorkers []*model.EngineerWorker

	rows, err := repo.db.QueryContext(ctx, `
		SELECT ew.id, ew.name, ew.surname, ew.patronymic, ew.age, ew.seniority, ew.building_organization_id, ep.id, ep.name 
		FROM building_site AS bs
        	JOIN construction_project AS cp ON bs.id = cp.building_site_id
        	JOIN engineer_team AS et ON cp.id = et.project_id
        	JOIN engineer_worker AS ew ON et.id = ew.team_id
        	JOIN engineer_position AS ep ON ew.position_id = ep.id
		WHERE building_site_id = $1
 		AND building_site_id != 0
	`, buildingSiteId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.EngineerWorker
		if err := rows.Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Seniority, &entity.BuildingOrganizationID, &entity.Position.ID, &entity.Position.Name); err != nil {
			return nil, err
		}
		engineerWorkers = append(engineerWorkers, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return engineerWorkers, nil
}

func (repo *engineerWorkerRepository) FindByConstructionManagement(ctx context.Context, constructionManagementId int) ([]*model.EngineerWorker, error) {
	var engineerWorkers []*model.EngineerWorker

	rows, err := repo.db.QueryContext(ctx, `
		SELECT ew.id, ew.name, ew.surname, ew.patronymic, ew.age, ew.seniority, ew.building_organization_id, ep.id, ep.name 
		FROM construction_management AS cm
        	JOIN building_site AS bs ON cm.id = bs.management_id
        	JOIN construction_project AS cp ON bs.id = cp.building_site_id
        	JOIN engineer_team AS et ON cp.id = et.project_id
        	JOIN engineer_worker AS ew ON et.id = ew.team_id
        	JOIN engineer_position AS ep ON ew.position_id = ep.id
		WHERE cm.id = $1
 		AND cm.id != 0
	`, constructionManagementId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.EngineerWorker
		if err := rows.Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Seniority, &entity.BuildingOrganizationID, &entity.Position.ID, &entity.Position.Name); err != nil {
			return nil, err
		}

		engineerWorkers = append(engineerWorkers, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return engineerWorkers, nil
}

func (repo *engineerWorkerRepository) FindAll(ctx context.Context) ([]*model.EngineerWorker, error) {
	var engineerWorkers []*model.EngineerWorker

	rows, err := repo.db.QueryContext(ctx, `
		SELECT ew.id, ew.name, ew.surname, ew.patronymic, ew.age, ew.seniority, ew.building_organization_id, ew.team_id, ep.id, ep.name
		FROM engineer_worker AS ew
		JOIN engineer_position AS ep ON ep.id = ew.position_id
		WHERE ew.id != 0
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.EngineerWorker
		err := rows.Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Seniority, &entity.BuildingOrganizationID, &entity.TeamID, &entity.Position.ID, &entity.Position.Name)
		if err != nil {
			return nil, err
		}

		engineerWorkers = append(engineerWorkers, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return engineerWorkers, nil
}
