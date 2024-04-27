package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EstimateHandlers struct {
	estimateService *service.EstimateService
}

func NewEstimateHandlers(service *service.EstimateService) *EstimateHandlers {
	return &EstimateHandlers{estimateService: service}
}

func (h *EstimateHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *EstimateHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *EstimateHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *EstimateHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *EstimateHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}

func (h *EstimateHandlers) GetExceededUsageMaterials(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	estimateID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	estimate, err := h.estimateService.GetExceededUsageMaterials(estimateID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(estimate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
