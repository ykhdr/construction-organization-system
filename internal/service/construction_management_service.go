package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionManagementService struct {
	constructionManagementRepository repository.ConstructionManagementRepository
	engineerWorkerRepository         repository.EngineerWorkerRepository
}

func NewConstructionManagementService(constructionManagementRepo repository.ConstructionManagementRepository, engineerWorkerRepo repository.EngineerWorkerRepository) *ConstructionManagementService {
	return &ConstructionManagementService{constructionManagementRepository: constructionManagementRepo, engineerWorkerRepository: engineerWorkerRepo}
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

func (s *ConstructionManagementService) GetEngineers(id int) ([]*model.EngineerWorker, error) {
	ctx := context.Background()

	engineers, err := s.engineerWorkerRepository.FindByConstructionManagement(ctx, id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting engineers by building site id")
		return nil, err
	}

	return engineers, nil
}
