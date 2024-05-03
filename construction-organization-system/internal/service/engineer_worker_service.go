package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/model"
	"context"
)

type EngineerWorkerService struct {
	engineerWorkerRepository repository.EngineerWorkerRepository
}

func NewEngineerWorkerService(repo repository.EngineerWorkerRepository) *EngineerWorkerService {
	return &EngineerWorkerService{engineerWorkerRepository: repo}
}

func (s *EngineerWorkerService) GetById(id int) (*model.EngineerWorker, error) {
	ctx := context.Background()
	worker, err := s.engineerWorkerRepository.Find(ctx, id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting engineer worker by id")
		return worker, err
	}

	return worker, nil
}

func (s *EngineerWorkerService) GetList() ([]*model.EngineerWorker, error) {
	workers, err := s.engineerWorkerRepository.FindAll(context.Background())
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on getting engineer worker list")
		return workers, err
	}
	return workers, nil
}
