package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type BuildingOrganizationHandlers struct {
	buildingOrganizationService *service.BuildingOrganizationService
}

func NewBuildingOrganizationHandlers(service *service.BuildingOrganizationService) *BuildingOrganizationHandlers {
	return &BuildingOrganizationHandlers{buildingOrganizationService: service}
}

func (h *BuildingOrganizationHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingOrganizationHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingOrganizationHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingOrganizationHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingOrganizationHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
