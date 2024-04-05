package service

import "construction-organization-system/internal/database/repository"

type EstimateService struct {
	estimateRepository repository.EstimateRepository
}

func NewEstimateService(repo repository.EstimateRepository) *EstimateService {
	return &EstimateService{estimateRepository: repo}
}
