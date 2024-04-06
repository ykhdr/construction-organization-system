package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/model"
	"context"
)

type BuildingSiteService struct {
	buildingSiteRepository   repository.BuildingSiteRepository
	engineerWorkerRepository repository.EngineerWorkerRepository
}

func NewBuildingSiteService(buildingSiteRepository repository.BuildingSiteRepository, engineerWorkerRepo repository.EngineerWorkerRepository) *BuildingSiteService {
	return &BuildingSiteService{buildingSiteRepository: buildingSiteRepository, engineerWorkerRepository: engineerWorkerRepo}
}

func (s *BuildingSiteService) GetById(id int) (*model.BuildingSite, error) {
	ctx := context.Background()
	buildingSite, err := s.buildingSiteRepository.Find(ctx, id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting building site by id")
		return nil, err
	}

	return buildingSite, nil
}

func (s *BuildingSiteService) GetList() ([]*model.BuildingSite, error) {
	ctx := context.Background()

	buildingSites, err := s.buildingSiteRepository.FindAll(ctx)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting building site list")
		return nil, err
	}

	return buildingSites, nil
}

func (s *BuildingSiteService) GetEngineers(id int) ([]*model.EngineerWorker, error) {
	ctx := context.Background()

	engineers, err := s.engineerWorkerRepository.FindByBuildingSite(ctx, id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting engineers by building site id")
		return nil, err
	}

	return engineers, nil
}
