package service

import "construction-organization-system/internal/database/repository"

type ConstructionProjectService struct {
	constructionProjectRepository repository.ConstructionProjectRepository
}

func NewConstructionProjectService(repo repository.ConstructionProjectRepository) *ConstructionProjectService {
	return &ConstructionProjectService{constructionProjectRepository: repo}
}
