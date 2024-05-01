package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type WorkScheduleService struct {
	workScheduleRepository repository.WorkScheduleRepository
}

func NewWorkScheduleService(repo repository.WorkScheduleRepository) *WorkScheduleService {
	return &WorkScheduleService{workScheduleRepository: repo}
}

func (s *WorkScheduleService) GetList() ([]*model.WorkSchedule, error) {
	schedules, err := s.workScheduleRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}

	return schedules, nil
}
