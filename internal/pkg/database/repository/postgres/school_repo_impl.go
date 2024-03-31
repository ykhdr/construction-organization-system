package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type schoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) repository.SchoolRepository {
	return &schoolRepository{db: db}
}

func (repo *schoolRepository) Save(ctx context.Context, entity model.School) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO school(name, building_site_id, classroom_count, floors) VALUES ($1, $2, $3, $4)",
		entity.Name, entity.BuildingSiteID, entity.ClassroomCount, entity.Floors).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *schoolRepository) Find(ctx context.Context, id int) (*model.School, error) {
	var entity model.School
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, building_site_id, classroom_count, floors FROM school WHERE id = $1", id).
		Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.ClassroomCount, &entity.Floors)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *schoolRepository) Update(ctx context.Context, entity model.School) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE school SET name = $1, building_site_id = $2, classroom_count = $3, floors = $4 WHERE id = $5",
		entity.Name, entity.BuildingSiteID, entity.ClassroomCount, entity.Floors, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *schoolRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM school WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
