package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type EmployeeHandlers struct {
	employeeService *service.EmployeeService
}

func NewEmployeeHandlers(service *service.EmployeeService) *EmployeeHandlers {
	return &EmployeeHandlers{employeeService: service}
}

func (h *EmployeeHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *EmployeeHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *EmployeeHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *EmployeeHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *EmployeeHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
