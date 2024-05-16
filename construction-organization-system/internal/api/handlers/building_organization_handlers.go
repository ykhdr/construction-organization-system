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
	vars := mux.Vars(r)

	organizationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error parsing organization id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	organization, err := h.buildingOrganizationService.Get(organizationID)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting organization")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(organization)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error encoding organization")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BuildingOrganizationHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingOrganizationHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingOrganizationHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingOrganizationHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	organizations, err := h.buildingOrganizationService.GetList()
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting organizations")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(organizations)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error encoding organizations")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
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
