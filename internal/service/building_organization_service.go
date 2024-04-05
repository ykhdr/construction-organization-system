package service

import (
	"construction-organization-system/internal/database/repository"
)

type BuildingOrganizationService struct {
	buildingOrganizationRepository repository.BuildingOrganizationRepository
}

func NewBuildingOrganizationService(repo repository.BuildingOrganizationRepository) *BuildingOrganizationService {
	return &BuildingOrganizationService{buildingOrganizationRepository: repo}
}
