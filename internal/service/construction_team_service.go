package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionTeamService struct {
	constructionTeamRepository   repository.ConstructionTeamRepository
	constructionWorkerRepository repository.ConstructionWorkerRepository
}

func (s *ConstructionTeamService) GetWorkers(teamId int) ([]*model.ConstructionWorker, error) {
	workers, err := s.constructionWorkerRepository.FindByTeam(context.Background(), teamId)
	if err != nil {
		return nil, err
	}
	return workers, nil
}

func NewConstructionTeamService(repo repository.ConstructionTeamRepository) *ConstructionTeamService {
	return &ConstructionTeamService{constructionTeamRepository: repo}
}
