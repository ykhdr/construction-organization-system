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

func (s *BuildingOrganizationService) GetList() ([]*model.BuildingOrganization, error) {
	organizations, err := s.buildingOrganizationRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func (s *BuildingOrganizationService) Get(id int) (*model.BuildingOrganization, error) {
	organization, err := s.buildingOrganizationRepository.Find(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return organization, nil
}
