package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionTeamService struct {
	constructionTeamRepository   repository.ConstructionTeamRepository
	constructionWorkerRepository repository.ConstructionWorkerRepository
	workTypeRepository           repository.WorkTypeRepository
}

func (s *ConstructionTeamService) GetWorkers(teamId int) ([]*model.ConstructionWorker, error) {
	workers, err := s.constructionWorkerRepository.FindByTeam(context.Background(), teamId)
	if err != nil {
		return nil, err
	}
	return workers, nil
}

func (s *ConstructionTeamService) GetWorkTypes(teamId int, startDate, endDate string) ([]*model.WorkType, error) {
	workTypes, err := s.workTypeRepository.FindByTeamWithPeriod(context.Background(), teamId, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return workTypes, nil
}

func (s *ConstructionTeamService) GetTeams(workTypeId int, startDate string, endDate string) ([]*model.ConstructionTeam, error) {
	teams, err := s.constructionTeamRepository.FindByWorkTypeWithPeriod(context.Background(), workTypeId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func NewConstructionTeamService(repo repository.ConstructionTeamRepository) *ConstructionTeamService {
	return &ConstructionTeamService{constructionTeamRepository: repo}
}
