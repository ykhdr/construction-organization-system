package service

import "construction-organization-system/internal/database/repository"

type EngineerTeamService struct {
	engineerTeamRepository repository.EngineerTeamRepository
}

func NewEngineerTeamService(repo repository.EngineerTeamRepository) *EngineerTeamService {
	return &EngineerTeamService{engineerTeamRepository: repo}
}
