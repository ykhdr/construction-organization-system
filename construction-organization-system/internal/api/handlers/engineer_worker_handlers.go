package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EngineerWorkerHandlers struct {
	engineerWorkerService *service.EngineerWorkerService
}

func NewEngineerWorkerHandlers(service *service.EngineerWorkerService) *EngineerWorkerHandlers {
	return &EngineerWorkerHandlers{engineerWorkerService: service}
}

func (h *EngineerWorkerHandlers) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	worker, err := h.engineerWorkerService.GetById(id)
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

func (h *EngineerWorkerHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerWorkerHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerWorkerHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerWorkerHandlers) GetList(w http.ResponseWriter, r *http.Request) {

	workers, err := h.engineerWorkerService.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
