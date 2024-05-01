package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"net/http"
)

type WorkScheduleHandlers struct {
	workScheduleService *service.WorkScheduleService
}

func NewWorkScheduleHandlers(service *service.WorkScheduleService) *WorkScheduleHandlers {
	return &WorkScheduleHandlers{workScheduleService: service}
}

func (h *WorkScheduleHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *WorkScheduleHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *WorkScheduleHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *WorkScheduleHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *WorkScheduleHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.workScheduleService.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(schedules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
