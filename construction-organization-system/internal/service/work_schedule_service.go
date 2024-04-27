package service

import "construction-organization-system/internal/database/repository"

type WorkScheduleService struct {
	workScheduleRepository repository.WorkScheduleRepository
}

func NewWorkScheduleService(repo repository.WorkScheduleRepository) *WorkScheduleService {
	return &WorkScheduleService{workScheduleRepository: repo}
}
