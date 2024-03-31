package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
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
	err := repo.db.QueryRowContext(ctx, "INSERT INTO engineer_worker(name, surname, patronymic, age, seniority, building_organization_id, skill_level, position_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID, entity.SkillLevel, entity.PositionID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *engineerWorkerRepository) Find(ctx context.Context, id int) (*model.EngineerWorker, error) {
	var entity model.EngineerWorker
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, surname, patronymic, age, seniority, building_organization_id, skill_level, position_id FROM engineer_worker WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Seniority, &entity.BuildingOrganizationID, &entity.SkillLevel, &entity.PositionID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *engineerWorkerRepository) Update(ctx context.Context, entity model.EngineerWorker) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE engineer_worker SET name = $1, surname = $2, patronymic = $3, age = $4, seniority = $5, building_organization_id = $6, skill_level = $7, team_id = $8, position_id = $9 WHERE id = $10",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID, entity.SkillLevel, entity.PositionID, entity.ID)
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
