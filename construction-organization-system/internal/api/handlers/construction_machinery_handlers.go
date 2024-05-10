package handlers

import (
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/service"
	"encoding/json"
	"net/http"
)

type ConstructionMachineryHandlers struct {
	constructionMachineryService *service.ConstructionMachineryService
}

func NewConstructionMachineryHandlers(service *service.ConstructionMachineryService) *ConstructionMachineryHandlers {
	return &ConstructionMachineryHandlers{constructionMachineryService: service}
}

func (h *ConstructionMachineryHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionMachineryHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionMachineryHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionMachineryHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionMachineryHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	machinery, err := h.constructionMachineryService.GetList()
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting machinery")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(machinery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
