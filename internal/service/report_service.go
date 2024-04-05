package service

import "construction-organization-system/internal/database/repository"

type ReportService struct {
	reportRepository repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) *ReportService {
	return &ReportService{reportRepository: repo}
}
