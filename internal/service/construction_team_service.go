package service

import "construction-organization-system/internal/database/repository"

type ConstructionTeamService struct {
	constructionTeamRepository repository.ConstructionTeamRepository
}

func NewConstructionTeamService(repo repository.ConstructionTeamRepository) *ConstructionTeamService {
	return &ConstructionTeamService{constructionTeamRepository: repo}
}
