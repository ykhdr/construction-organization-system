package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type BuildingOrganizationService struct {
	buildingOrganizationRepository repository.BuildingOrganizationRepository
	constructionProjectRepository  repository.ConstructionProjectRepository
}

func NewBuildingOrganizationService(repo repository.BuildingOrganizationRepository) *BuildingOrganizationService {
	return &BuildingOrganizationService{buildingOrganizationRepository: repo}
}

func (s *BuildingOrganizationService) GetProjects(organizationID int) ([]*model.ConstructionProject, error) {
	projects, err := s.constructionProjectRepository.FindByOrganization(context.Background(), organizationID)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
