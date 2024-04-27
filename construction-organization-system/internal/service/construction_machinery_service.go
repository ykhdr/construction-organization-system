package service

import "construction-organization-system/internal/database/repository"

type ConstructionMachineryService struct {
	constructionMachineryRepository repository.ConstructionMachineryRepository
}

func NewConstructionMachineryService(repo repository.ConstructionMachineryRepository) *ConstructionMachineryService {
	return &ConstructionMachineryService{constructionMachineryRepository: repo}
}
