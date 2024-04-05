package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
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
