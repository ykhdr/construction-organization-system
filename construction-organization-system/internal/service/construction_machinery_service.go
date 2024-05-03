package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionMachineryService struct {
	constructionMachineryRepository repository.ConstructionMachineryRepository
}

func (s *ConstructionMachineryService) GetList() ([]*model.ConstructionMachinery, error) {
	machines, err := s.constructionMachineryRepository.FindAll(context.Background())

	if err != nil {
		return nil, err
	}

	return machines, nil
}

func NewConstructionMachineryService(repo repository.ConstructionMachineryRepository) *ConstructionMachineryService {
	return &ConstructionMachineryService{constructionMachineryRepository: repo}
}
