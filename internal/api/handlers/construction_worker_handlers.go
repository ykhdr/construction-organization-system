package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type ConstructionWorkerHandlers struct {
	constructionWorkerService *service.ConstructionWorkerService
}

func NewConstructionWorkerHandlers(service *service.ConstructionWorkerService) *ConstructionWorkerHandlers {
	return &ConstructionWorkerHandlers{constructionWorkerService: service}
}

func (h *ConstructionWorkerHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionWorkerHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionWorkerHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionWorkerHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionWorkerHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
