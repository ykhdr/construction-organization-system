package handlers

import (
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ConstructionManagementHandlers struct {
	constructionManagementService *service.ConstructionManagementService
	engineerWorkerService         *service.EngineerWorkerService
}

func NewConstructionManagementHandlers(managementService *service.ConstructionManagementService, engineerWorkerService *service.EngineerWorkerService) *ConstructionManagementHandlers {
	return &ConstructionManagementHandlers{
		constructionManagementService: managementService,
		engineerWorkerService:         engineerWorkerService,
	}
}

func (h *ConstructionManagementHandlers) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get construction management id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	management, err := h.constructionManagementService.GetById(id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting construction management")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(management)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionManagementHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionManagementHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionManagementHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionManagementHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	managements, err := h.constructionManagementService.GetList()
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting construction management")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(managements)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *ConstructionManagementHandlers) GetManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	managementId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get construction management id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	management, err := h.constructionManagementService.GetById(managementId)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting construction management")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	worker, err := h.engineerWorkerService.GetById(management.ManagerID)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting worker")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(worker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionManagementHandlers) GetEngineers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	managementId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get construction management id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	engineers, err := h.constructionManagementService.GetEngineers(managementId)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting engineers")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(engineers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionManagementHandlers) GetProjects(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	managementId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get construction management id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	projects, err := h.constructionManagementService.GetProjects(managementId)
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

func (h *ConstructionManagementHandlers) GetMachines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	managementId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get construction management id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	machines, err := h.constructionManagementService.GetMachines(managementId)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting machines")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(machines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
