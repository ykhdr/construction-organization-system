package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"net/http"
)

type WorkTypeHandlers struct {
	workTypeService *service.WorkTypeService
}

func NewWorkTypeHandlers(service *service.WorkTypeService) *WorkTypeHandlers {
	return &WorkTypeHandlers{workTypeService: service}
}

func (h *WorkTypeHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	workTypes, err := h.workTypeService.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workTypes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
