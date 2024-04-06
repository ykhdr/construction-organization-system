package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionManagementService struct {
	constructionManagementRepository repository.ConstructionManagementRepository
}

func NewConstructionManagementService(repo repository.ConstructionManagementRepository) *ConstructionManagementService {
	return &ConstructionManagementService{constructionManagementRepository: repo}
}

func (s *ConstructionManagementService) GetById(id int) (*model.ConstructionManagement, error) {
	ctx := context.Background()

	management, err := s.constructionManagementRepository.Find(ctx, id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting construction management by id")
		return nil, err
	}

	return management, nil
}

func (s *ConstructionManagementService) GetList() ([]*model.ConstructionManagement, error) {
	ctx := context.Background()

	managements, err := s.constructionManagementRepository.FindAll(ctx)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting construction management list")
		return nil, err
	}

	return managements, nil
}
