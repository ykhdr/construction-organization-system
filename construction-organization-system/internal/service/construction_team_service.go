package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"encoding/json"
	"io"
)

type ConstructionTeamService struct {
	constructionTeamRepository   repository.ConstructionTeamRepository
	constructionWorkerRepository repository.ConstructionWorkerRepository
	workTypeRepository           repository.WorkTypeRepository
}

func NewConstructionTeamService(constructionTeamRepository repository.ConstructionTeamRepository, constructionWorkerRepository repository.ConstructionWorkerRepository, workTypeRepository repository.WorkTypeRepository) *ConstructionTeamService {
	return &ConstructionTeamService{constructionTeamRepository: constructionTeamRepository, constructionWorkerRepository: constructionWorkerRepository, workTypeRepository: workTypeRepository}
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

func (s *ConstructionTeamService) Get(teamId int) (*model.ConstructionTeam, error) {
	team, err := s.constructionTeamRepository.Find(context.Background(), teamId)

	if err != nil {
		return nil, err
	}
	return team, nil
}

func (s *ConstructionTeamService) Create(body io.ReadCloser) error {
	var team model.ConstructionTeam

	err := json.NewDecoder(body).Decode(&team)
	if err != nil {
		return err
	}

	_, err = s.constructionTeamRepository.Save(context.Background(), team)
	if err != nil {
		return err
	}
	return nil
}

func (s *ConstructionTeamService) Delete(id int) error {
	return s.constructionTeamRepository.Delete(context.Background(), id)
}
