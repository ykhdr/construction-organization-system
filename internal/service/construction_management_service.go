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
	projectRepository                repository.ConstructionProjectRepository
	constructionMachineryRepository  repository.ConstructionMachineryRepository
}

func NewConstructionManagementService(constructionManagementRepo repository.ConstructionManagementRepository, engineerWorkerRepo repository.EngineerWorkerRepository) *ConstructionManagementService {
	return &ConstructionManagementService{constructionManagementRepository: constructionManagementRepo, engineerWorkerRepository: engineerWorkerRepo}
}

func (s *ConstructionManagementService) GetById(id int) (*model.ConstructionManagement, error) {
	management, err := s.constructionManagementRepository.Find(context.Background(), id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting construction management by id")
		return nil, err
	}

	return management, nil
}

func (s *ConstructionManagementService) GetList() ([]*model.ConstructionManagement, error) {
	managements, err := s.constructionManagementRepository.FindAll(context.Background())
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting construction management list")
		return nil, err
	}

	return managements, nil
}

func (s *ConstructionManagementService) GetEngineers(id int) ([]*model.EngineerWorker, error) {
	engineers, err := s.engineerWorkerRepository.FindByConstructionManagement(context.Background(), id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting engineers by building site id")
		return nil, err
	}

	return engineers, nil
}

func (s *ConstructionManagementService) GetProjects(id int) ([]*model.ConstructionProject, error) {
	projects, err := s.projectRepository.FindByConstructionManagement(context.Background(), id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting projects by building site id")
		return nil, err
	}

	return projects, nil
}

func (s *ConstructionManagementService) GetMachines(id int) ([]*model.ConstructionMachinery, error) {
	machines, err := s.constructionMachineryRepository.GetByManagement(context.Background(), id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting machines by building site id")
		return nil, err
	}

	return machines, nil
}
