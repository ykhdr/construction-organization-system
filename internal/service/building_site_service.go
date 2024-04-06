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
	projectRepository        repository.ConstructionProjectRepository
}

func NewBuildingSiteService(buildingSiteRepository repository.BuildingSiteRepository, engineerWorkerRepo repository.EngineerWorkerRepository) *BuildingSiteService {
	return &BuildingSiteService{buildingSiteRepository: buildingSiteRepository, engineerWorkerRepository: engineerWorkerRepo}
}

func (s *BuildingSiteService) GetById(id int) (*model.BuildingSite, error) {
	buildingSite, err := s.buildingSiteRepository.Find(context.Background(), id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting building site by id")
		return nil, err
	}

	return buildingSite, nil
}

func (s *BuildingSiteService) GetList() ([]*model.BuildingSite, error) {
	buildingSites, err := s.buildingSiteRepository.FindAll(context.Background())
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting building site list")
		return nil, err
	}

	return buildingSites, nil
}

func (s *BuildingSiteService) GetEngineers(id int) ([]*model.EngineerWorker, error) {
	engineers, err := s.engineerWorkerRepository.FindByBuildingSite(context.Background(), id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting engineers by building site id")
		return nil, err
	}

	return engineers, nil
}

func (s *BuildingSiteService) GetProjects(id int) ([]*model.ConstructionProject, error) {
	projects, err := s.projectRepository.FindByBuildingSite(context.Background(), id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting projects by building site id")
		return nil, err
	}

	return projects, nil
}
