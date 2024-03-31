package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type constructionWorkerRepository struct {
	db *sql.DB
}

func NewConstructionWorkerRepository(db *sql.DB) repository.ConstructionWorkerRepository {
	return &constructionWorkerRepository{db: db}
}

func (repo *constructionWorkerRepository) Save(ctx context.Context, entity model.ConstructionWorker) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_worker(name, surname, patronymic, age, seniority, building_organization_id, is_shift_worker, team_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID, entity.IsShiftWorker, entity.TeamID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionWorkerRepository) Find(ctx context.Context, id int) (*model.ConstructionWorker, error) {
	var entity model.ConstructionWorker
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, surname, patronymic, age, seniority, building_organization_id, is_shift_worker, team_id FROM construction_worker WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Seniority, &entity.BuildingOrganizationID, &entity.IsShiftWorker, &entity.TeamID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionWorkerRepository) Update(ctx context.Context, entity model.ConstructionWorker) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_worker SET name = $1, surname = $2, patronymic = $3, age = $4, seniority = $5, building_organization_id = $6, is_shift_worker = $7, team_id = $8 WHERE id = $9",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID, entity.IsShiftWorker, entity.TeamID, entity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *constructionWorkerRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_worker WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
