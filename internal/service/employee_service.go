package service

import "construction-organization-system/internal/database/repository"

type EmployeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepository: repo}
}
