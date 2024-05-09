package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type EngineerTeamService struct {
	engineerTeamRepository repository.EngineerTeamRepository
}

func (s *EngineerTeamService) Get(id int) (*model.EngineerTeam, error) {
	team, err := s.engineerTeamRepository.Find(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (s *EngineerTeamService) GetList() ([]*model.EngineerTeam, error) {
	teams, err := s.engineerTeamRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func NewEngineerTeamService(repo repository.EngineerTeamRepository) *EngineerTeamService {
	return &EngineerTeamService{engineerTeamRepository: repo}
}
