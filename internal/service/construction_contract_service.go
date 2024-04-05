package service

import "construction-organization-system/internal/database/repository"

type ConstructionContractService struct {
	constructionContractRepository repository.ConstructionContractRepository
}

func NewConstructionContractService(repo repository.ConstructionContractRepository) *ConstructionContractService {
	return &ConstructionContractService{constructionContractRepository: repo}
}
