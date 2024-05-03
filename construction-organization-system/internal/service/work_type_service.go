package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type WorkTypeService struct {
	workTypeRepository repository.WorkTypeRepository
}

func NewWorkTypeService(repo repository.WorkTypeRepository) *WorkTypeService {
	return &WorkTypeService{workTypeRepository: repo}
}

func (s *WorkTypeService) GetList() ([]*model.WorkType, error) {
	workTypes, err := s.workTypeRepository.FindAll(context.Background())

	if err != nil {
		return nil, err
	}

	return workTypes, nil
}
