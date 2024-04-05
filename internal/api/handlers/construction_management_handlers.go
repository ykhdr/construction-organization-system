package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ConstructionManagementHandlers struct {
	constructionManagementService *service.ConstructionManagementService
}

func NewConstructionManagementHandlers(service *service.ConstructionManagementService) *ConstructionManagementHandlers {
	return &ConstructionManagementHandlers{constructionManagementService: service}
}

func (h *ConstructionManagementHandlers) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	management, err := h.constructionManagementService.GetById(id)
	if err != nil {
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

}
