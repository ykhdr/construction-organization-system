package service

import "construction-organization-system/internal/database/repository"

type BuildingSiteService struct {
	buildingSiteRepository repository.BuildingSiteRepository
}

func NewBuildingSiteService(repo repository.BuildingSiteRepository) *BuildingSiteService {
	return &BuildingSiteService{buildingSiteRepository: repo}
}
