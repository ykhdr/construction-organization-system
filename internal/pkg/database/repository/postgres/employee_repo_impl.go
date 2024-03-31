package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) repository.EmployeeRepository {
	return &employeeRepository{db: db}
}

func (repo *employeeRepository) Save(ctx context.Context, entity model.Employee) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO employee(name, surname, patronymic, age, seniority, building_organization_id) VALUES ($1, $2, $3, $4, $5, $6)",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *employeeRepository) Find(ctx context.Context, id int) (*model.Employee, error) {
	var entity model.Employee
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, surname, patronymic, age, seniority, building_organization_id FROM employee WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Seniority, &entity.BuildingOrganizationID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *employeeRepository) Update(ctx context.Context, entity model.Employee) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE employee SET name = $1, surname = $2, patronymic = $3, age = $4, seniority = $5, building_organization_id = $6 WHERE id = $7",
		entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Seniority, entity.BuildingOrganizationID, entity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *employeeRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM employee WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
