package service

import "construction-organization-system/internal/database/repository"

type ApartmentHouseService struct {
	apartmentHouseRepository repository.ApartmentHouseRepository
}

func NewApartmentHouseService(repo repository.ApartmentHouseRepository) *ApartmentHouseService {
	return &ApartmentHouseService{apartmentHouseRepository: repo}
}
