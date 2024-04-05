package postgres

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"database/sql"
)

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) repository.ReportRepository {
	return &reportRepository{db: db}
}

func (repo *reportRepository) Save(ctx context.Context, entity model.Report) (int, error) {
	var newId int
	err := repo.db.QueryRowContext(ctx, "INSERT INTO report(project_id, report_creation_date, report_file) VALUES ($1, $2, $3)",
		entity.ProjectID, entity.ReportCreationDate, entity.ReportFile).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func (repo *reportRepository) Find(ctx context.Context, id int) (*model.Report, error) {
	var entity model.Report
	err := repo.db.QueryRowContext(ctx, "SELECT id, project_id, report_creation_date, report_file FROM report WHERE id = $1", id).
		Scan(&entity.ID, &entity.ProjectID, &entity.ReportCreationDate, &entity.ReportFile)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *reportRepository) Update(ctx context.Context, entity model.Report) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE report SET project_id = $1, report_creation_date = $2, report_file = $3 WHERE id = $4",
		entity.ProjectID, entity.ReportCreationDate, entity.ReportFile, entity.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *reportRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM report WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
