package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/model"
	"context"
)

type BuildingSiteService struct {
	buildingSiteRepository repository.BuildingSiteRepository
}

func NewBuildingSiteService(repo repository.BuildingSiteRepository) *BuildingSiteService {
	return &BuildingSiteService{buildingSiteRepository: repo}
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
