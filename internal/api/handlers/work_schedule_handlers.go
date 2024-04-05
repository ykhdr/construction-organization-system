package handlers

import (
	"construction-organization-system/internal/service"
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

}
