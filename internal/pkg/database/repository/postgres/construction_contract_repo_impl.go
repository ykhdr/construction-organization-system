package postgres

import (
	"construction-organization-system/internal/pkg/database/repository"
	"construction-organization-system/internal/pkg/model"
	"context"
	"database/sql"
)

type constructionContractRepository struct {
	db *sql.DB
}

func NewConstructionContractRepository(db *sql.DB) repository.ConstructionContractRepository {
	return &constructionContractRepository{db: db}
}

func (repo *constructionContractRepository) Save(ctx context.Context, entity model.ConstructionContract) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO construction_contract(building_organization_id, customer_organization_id, name, project_id, signing_date) VALUES ($1,$2, $3, $4, $5)",
		entity.BuildingOrganizationID, entity.CustomerOrganizationID, entity.Name, entity.ProjectID, entity.SigningDate).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *constructionContractRepository) Find(ctx context.Context, id int) (*model.ConstructionContract, error) {
	var entity model.ConstructionContract
	err := repo.db.QueryRowContext(ctx, "SELECT id, building_organization_id, customer_organization_id, name, project_id, signing_date FROM construction_contract ah WHERE id = $1 AND id != 0", id).
		Scan(&entity.ID, &entity.BuildingOrganizationID, &entity.CustomerOrganizationID, &entity.Name, &entity.ProjectID, &entity.SigningDate)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *constructionContractRepository) Update(ctx context.Context, entity model.ConstructionContract) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE construction_contract SET building_organization_id = $1, customer_organization_id = $2, name = $3, project_id = $4, signing_date = $5 WHERE id = $6",
		entity.BuildingOrganizationID, entity.CustomerOrganizationID, entity.Name, entity.ProjectID, entity.SigningDate, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *constructionContractRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM construction_contract WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
