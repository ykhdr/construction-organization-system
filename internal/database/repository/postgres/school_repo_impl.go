package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/model"
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
	tr, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	var newId int
	err = tr.QueryRowContext(ctx, `
	INSERT INTO construction_project(name, building_site_id, type) 
	VALUES ($1, $2, $3) 
	RETURNING id;
	`, entity.Name, entity.BuildingSiteID, "SCHOOL").Scan(&newId)
	if err != nil {
		return 0, err
	}

	_, err = tr.ExecContext(ctx, `
	INSERT INTO school(project_id, classroom_count, floors)
	VALUES ($1, $2, $3)
	`, newId, entity.ClassroomCount, entity.Floors)

	if err != nil {
		log.Logger.WithError(err).Error("Error inserting school")
		tr.Rollback()
		return -1, err
	}

	if err := tr.Commit(); err != nil {
		return -1, err
	}

	return newId, nil
}

func (repo *schoolRepository) Find(ctx context.Context, id int) (*model.School, error) {
	var entity model.School

	err := repo.db.QueryRowContext(ctx, `
		SELECT id, name, building_site_id, type, classroom_count, floors 
		FROM school
			JOIN construction_project cp on school.project_id = cp.id
		WHERE id = $1`, id).
		Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.ProjectType, &entity.ClassroomCount, &entity.Floors)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *schoolRepository) Update(ctx context.Context, entity model.School) error {
	tr, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = tr.ExecContext(ctx, `
		UPDATE construction_project
		SET name = $1, 
		    building_site_id = $2
		WHERE id = $3;
	`, entity.Name, entity.BuildingSiteID, entity.ID)
	if err != nil {
		return err
	}

	_, err = tr.ExecContext(ctx, `
	UPDATE school 
	SET classroom_count = $1, 
	    floors = $2 
	WHERE project_id = $3`,
		entity.ClassroomCount, entity.Floors, entity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *schoolRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM school WHERE project_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *schoolRepository) FindAll(ctx context.Context) ([]*model.School, error) {
	var entities []*model.School

	rows, err := repo.db.QueryContext(ctx, `
		SELECT id, name, building_site_id, type, classroom_count, floors 
		FROM school
			JOIN construction_project cp on school.project_id = cp.id
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.School
		err = rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.ProjectType, &entity.ClassroomCount, &entity.Floors)
		if err != nil {
			return nil, err
		}
		entities = append(entities, &entity)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
