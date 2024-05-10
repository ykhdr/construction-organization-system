package handlers

import (
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (h *BuildingOrganizationHandlers) GetProjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	organizationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get organization id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	projects, err := h.buildingOrganizationService.GetProjects(organizationID)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting projects")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
