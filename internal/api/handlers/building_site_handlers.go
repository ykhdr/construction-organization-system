package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BuildingSiteHandlers struct {
	buildingSiteService   *service.BuildingSiteService
	engineerWorkerService *service.EngineerWorkerService
}

func NewBuildingSiteHandlers(buildingSiteService *service.BuildingSiteService, engineerWorkerService *service.EngineerWorkerService) *BuildingSiteHandlers {
	return &BuildingSiteHandlers{buildingSiteService: buildingSiteService, engineerWorkerService: engineerWorkerService}
}

func (h *BuildingSiteHandlers) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["site-id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buildingSite, err := h.buildingSiteService.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(buildingSite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *BuildingSiteHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	buildingSites, err := h.buildingSiteService.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(buildingSites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *BuildingSiteHandlers) GetManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	siteId, err := strconv.Atoi(vars["site-id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buildingSite, err := h.buildingSiteService.GetById(siteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	worker, err := h.engineerWorkerService.GetById(buildingSite.ManagerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(worker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
