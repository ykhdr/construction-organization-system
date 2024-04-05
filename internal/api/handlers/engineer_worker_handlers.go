package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type EngineerWorkerHandlers struct {
	engineerWorkerService *service.EngineerWorkerService
}

func NewEngineerWorkerHandlers(service *service.EngineerWorkerService) *EngineerWorkerHandlers {
	return &EngineerWorkerHandlers{engineerWorkerService: service}
}

func (h *EngineerWorkerHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerWorkerHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerWorkerHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerWorkerHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerWorkerHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
