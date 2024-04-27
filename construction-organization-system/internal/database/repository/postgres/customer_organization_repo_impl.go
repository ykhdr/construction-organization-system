package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type customerOrganizationRepository struct {
	db *sql.DB
}

func NewCustomerOrganizationRepository(db *sql.DB) repository.CustomerOrganizationRepository {
	return &customerOrganizationRepository{db: db}
}

func (repo *customerOrganizationRepository) Save(ctx context.Context, entity model.CustomerOrganization) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO customer_organization(name) VALUES ($1)",
		entity.Name).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *customerOrganizationRepository) Find(ctx context.Context, id int) (*model.CustomerOrganization, error) {
	var entity model.CustomerOrganization
	err := repo.db.QueryRowContext(ctx, "SELECT id, name FROM customer_organization WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *customerOrganizationRepository) Update(ctx context.Context, entity model.CustomerOrganization) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE customer_organization SET name = $1 WHERE id = $2",
		entity.Name, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customerOrganizationRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM customer_organization WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
