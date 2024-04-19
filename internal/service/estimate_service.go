package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
)

type EstimateService struct {
	estimateRepository repository.EstimateRepository
	materialRepository repository.MaterialRepository
}

func NewEstimateService(repo repository.EstimateRepository) *EstimateService {
	return &EstimateService{estimateRepository: repo}
}

func (s *EstimateService) GetExceededUsageMaterials(estimateID int) ([]*model.Material, error) {
	materials, err := s.materialRepository.FindByEstimateWithExceededUsage(nil, estimateID)
	if err != nil {
		return nil, err
	}
	return materials, nil
}
