package service

import "construction-organization-system/internal/database/repository"

type CustomerOrganizationService struct {
	customerOrganizationRepository repository.CustomerOrganizationRepository
}

func NewCustomerOrganizationService(repo repository.CustomerOrganizationRepository) *CustomerOrganizationService {
	return &CustomerOrganizationService{customerOrganizationRepository: repo}
}
