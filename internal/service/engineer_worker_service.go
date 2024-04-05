package service

import "construction-organization-system/internal/database/repository"

type EngineerWorkerService struct {
	engineerWorkerRepository repository.EngineerWorkerRepository
}

func NewEngineerWorkerService(repo repository.EngineerWorkerRepository) *EngineerWorkerService {
	return &EngineerWorkerService{engineerWorkerRepository: repo}
}
