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
	tr, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	var newId int
	err = tr.QueryRowContext(ctx, `
	INSERT INTO construction_project(name, building_site_id, type) 
	VALUES ($1, $2, $3) 
	RETURNING id;
	`, entity.Name, entity.BuildingSiteID, "BRIDGE").Scan(&newId)
	if err != nil {
		return 0, err
	}

	_, err = tr.ExecContext(ctx, `
	INSERT INTO bridge(project_id, span, width, traffic_lanes_number)
	VALUES ($1, $2, $3, $4)
	`, newId, entity.Span, entity.Width, entity.TrafficLanesNumber)

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

func (repo *bridgeRepository) Find(ctx context.Context, id int) (*model.Bridge, error) {
	var entity model.Bridge

	err := repo.db.QueryRowContext(ctx, `
		SELECT id, name, building_site_id, type, span, width, traffic_lanes_number
		FROM bridge
			JOIN construction_project cp on bridge.project_id = cp.id
		WHERE id = $1`, id).
		Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.ProjectType, &entity.Span, &entity.Width, &entity.TrafficLanesNumber)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *bridgeRepository) Update(ctx context.Context, entity model.Bridge) error {
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
	UPDATE bridge 
	SET  span = $1,
		width = $2,
		traffic_lanes_number = $3
	WHERE project_id = $4`,
		entity.Span, entity.Width, entity.TrafficLanesNumber, entity.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *bridgeRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM bridge WHERE project_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *bridgeRepository) FindAll(ctx context.Context) ([]*model.Bridge, error) {
	var entities []*model.Bridge

	rows, err := repo.db.QueryContext(ctx, `
		SELECT id, name, building_site_id, type, span, width, traffic_lanes_number 
		FROM bridge
			JOIN construction_project cp on bridge.project_id = cp.id
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity model.Bridge
		err = rows.Scan(&entity.ID, &entity.Name, &entity.BuildingSiteID, &entity.ProjectType, &entity.Span, &entity.Width, &entity.TrafficLanesNumber)
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
