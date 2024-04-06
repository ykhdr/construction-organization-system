package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionProjectService struct {
	constructionProjectRepository repository.ConstructionProjectRepository
	workScheduleRepository        repository.WorkScheduleRepository
}

func NewConstructionProjectService(constructionProjectRepo repository.ConstructionProjectRepository, workScheduleRepo repository.WorkScheduleRepository) *ConstructionProjectService {
	return &ConstructionProjectService{constructionProjectRepository: constructionProjectRepo, workScheduleRepository: workScheduleRepo}
}

func (s *ConstructionProjectService) GetWorkSchedules(projectID int) ([]*model.WorkSchedule, error) {
	workSchedules, err := s.workScheduleRepository.FindByProject(context.Background(), projectID)
	if err != nil {
		return nil, err
	}
	return workSchedules, nil
}
