package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
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
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO apartment_house(name, building_site_id, floors) VALUES ($1,$2,$3)",
		entity.Name, entity.BuildingSiteID, entity.Floors).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (repo *apartmentHouseRepository) Find(ctx context.Context, id int) (*model.ApartmentHouse, error) {
	var entity model.ApartmentHouse
	err := repo.db.QueryRowContext(ctx, "SELECT id, name,building_site_id, floors FROM apartment_house ah WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.Floors)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (repo *apartmentHouseRepository) Update(ctx context.Context, entity model.ApartmentHouse) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE apartment_house SET name = $1, building_site_id = $2, floors = $3 WHERE id = $4",
		entity.Name, entity.BuildingSiteID, entity.Floors, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *apartmentHouseRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM apartment_house WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
