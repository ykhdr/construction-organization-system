package service

import "construction-organization-system/internal/database/repository"

type ConstructionManagementService struct {
	constructionManagementRepository repository.ConstructionManagementRepository
}

func NewConstructionManagementService(repo repository.ConstructionManagementRepository) *ConstructionManagementService {
	return &ConstructionManagementService{constructionManagementRepository: repo}
}
