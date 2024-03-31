package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type buildingOrganizationRepository struct {
	db *sql.DB
}

func NewBuildingOrganizationRepository(db *sql.DB) repository.BuildingOrganizationRepository {
	return &buildingOrganizationRepository{db: db}
}

func (repo *buildingOrganizationRepository) Save(ctx context.Context, entity model.BuildingOrganization) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO building_organization(name) VALUES ($1)",
		entity.Name).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *buildingOrganizationRepository) Find(ctx context.Context, id int) (*model.BuildingOrganization, error) {
	var entity model.BuildingOrganization
	err := repo.db.QueryRowContext(ctx, "SELECT id, name FROM building_organization WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *buildingOrganizationRepository) Update(ctx context.Context, entity model.BuildingOrganization) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE building_organization SET name = $1 WHERE id = $2",
		entity.Name, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *buildingOrganizationRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM building_organization WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
