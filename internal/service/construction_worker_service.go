package service

import "construction-organization-system/internal/database/repository"

type ConstructionWorkerService struct {
	constructionWorkerRepository repository.ConstructionWorkerRepository
}

func NewConstructionWorkerService(repo repository.ConstructionWorkerRepository) *ConstructionWorkerService {
	return &ConstructionWorkerService{constructionWorkerRepository: repo}
}
