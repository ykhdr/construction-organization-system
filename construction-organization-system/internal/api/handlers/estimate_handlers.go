package handlers

import (
	"construction-organization-system/internal/log"
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
		log.Logger.WithError(err).Errorln("Error get estimate id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	materials, err := h.estimateService.GetExceededUsageMaterials(estimateID)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting exceeded usage materials")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(materials)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting exceeded usage materials")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
