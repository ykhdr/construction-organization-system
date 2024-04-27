package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type buildingSiteRepository struct {
	db *sql.DB
}

func NewBuildingSiteRepository(db *sql.DB) repository.BuildingSiteRepository {
	return &buildingSiteRepository{db: db}
}

func (repo *buildingSiteRepository) Save(ctx context.Context, entity model.BuildingSite) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO building_site(address, management_id, manager_id) VALUES ($1, $2, $3)",
		entity.Address, entity.ManagementID, entity.ManagerID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *buildingSiteRepository) Find(ctx context.Context, id int) (*model.BuildingSite, error) {
	var entity model.BuildingSite
	err := repo.db.QueryRowContext(ctx, "SELECT id, address, management_id, manager_id FROM building_site WHERE id = $1", id).
		Scan(&entity.ID, &entity.Address, &entity.ManagementID, &entity.ManagerID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *buildingSiteRepository) Update(ctx context.Context, entity model.BuildingSite) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE building_site SET address = $1, management_id = $2, manager_id = $3 WHERE id = $4",
		entity.Address, entity.ManagementID, entity.ManagerID, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *buildingSiteRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM building_site WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *buildingSiteRepository) FindAll(ctx context.Context) ([]*model.BuildingSite, error) {
	var buildingSites []*model.BuildingSite
	rows, err := repo.db.QueryContext(ctx, "SELECT id, address, management_id, manager_id FROM building_site WHERE id != 0")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var entity model.BuildingSite
		err = rows.Scan(&entity.ID, &entity.Address, &entity.ManagementID, &entity.ManagerID)
		if err != nil {
			return nil, err
		}
		buildingSites = append(buildingSites, &entity)
	}
	return buildingSites, nil
}
