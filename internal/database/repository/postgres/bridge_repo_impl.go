package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type bridgeRepository struct {
	db *sql.DB
}

func NewBridgeRepository(db *sql.DB) repository.BridgeRepository {
	return &bridgeRepository{db: db}
}

func (repo *bridgeRepository) Save(ctx context.Context, entity model.Bridge) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO bridge(name, building_site_id, span, width, traffic_lanes_number) VALUES ($1,$2,$3,$4,$5)",
		entity.Name, entity.BuildingSiteID, entity.Span, entity.Width, entity.TrafficLanesNumber).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (repo *bridgeRepository) Find(ctx context.Context, id int) (*model.Bridge, error) {
	var entity model.Bridge
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, building_site_id, span, width, traffic_lanes_number  FROM bridge ah WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.Span, &entity.Width, &entity.TrafficLanesNumber)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (repo *bridgeRepository) Update(ctx context.Context, entity model.Bridge) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE bridge SET name = $1, building_site_id = $2, span = $3, width = $4, traffic_lanes_number = $5 WHERE id = $6",
		entity.Name, entity.BuildingSiteID, entity.Span, entity.Width, entity.TrafficLanesNumber, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *bridgeRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM bridge WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
