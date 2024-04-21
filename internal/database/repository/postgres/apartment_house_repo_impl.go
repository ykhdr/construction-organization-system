package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type apartmentHouseRepository struct {
	db *sql.DB
}

func NewApartmentHouseRepository(db *sql.DB) repository.ApartmentHouseRepository {
	return &apartmentHouseRepository{db: db}
}

func (repo *apartmentHouseRepository) Save(ctx context.Context, entity model.ApartmentHouse) (int, error) {
	tr, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	var newId int
	err = tr.QueryRowContext(ctx, `
	INSERT INTO construction_project(name, building_site_id, type) 
	VALUES ($1, $2, $3) 
	RETURNING id;
	`, entity.Name, entity.BuildingSiteID, "APARTMENT_HOUSE").Scan(&newId)
	if err != nil {
		return 0, err
	}

	_, err = tr.ExecContext(ctx, `
	INSERT INTO apartment_house(project_id, floors)
	VALUES ($1, $2)
	`, newId, entity.Floors)

	if err != nil {
		err := tr.Rollback()
		if err != nil {
			return -1, err
		}
		return -1, err
	}

	if err := tr.Commit(); err != nil {
		return -1, err
	}

	return newId, nil
}

func (repo *apartmentHouseRepository) Find(ctx context.Context, id int) (*model.ApartmentHouse, error) {
	var entity model.ApartmentHouse

	err := repo.db.QueryRowContext(ctx, `
		SELECT id, name, building_site_id, type, floors 
		FROM school
			JOIN construction_project cp on school.project_id = cp.id
		WHERE id = $1`, id).
		Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.ProjectType, &entity.Floors)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *apartmentHouseRepository) Update(ctx context.Context, entity model.ApartmentHouse) error {
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
	SET  floors = $1
	WHERE project_id = $2`,
		entity.Floors, entity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *apartmentHouseRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM apartment_house WHERE project_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *apartmentHouseRepository) FindAll(ctx context.Context) ([]*model.ApartmentHouse, error) {
	var entities []*model.ApartmentHouse

	rows, err := repo.db.QueryContext(ctx, `
		SELECT id, name, building_site_id, type, floors 
		FROM school
			JOIN construction_project cp on school.project_id = cp.id
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.ApartmentHouse
		err = rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.ProjectType, &entity.Floors)
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
