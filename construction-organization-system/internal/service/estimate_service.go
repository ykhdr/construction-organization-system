package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type EstimateService struct {
	estimateRepository repository.EstimateRepository
	materialRepository repository.MaterialRepository
}

func NewEstimateService(estimateRepository repository.EstimateRepository, materialRepository repository.MaterialRepository) *EstimateService {
	return &EstimateService{estimateRepository: estimateRepository, materialRepository: materialRepository}
}

func (s *EstimateService) GetExceededUsageMaterials(estimateID int) ([]*model.Material, error) {
	materials, err := s.materialRepository.FindByEstimateWithExceededUsage(context.Background(), estimateID)
	if err != nil {
		return nil, err
	}
	return materials, nil
}
