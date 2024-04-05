package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type ConstructionManagementHandlers struct {
	constructionManagementService *service.ConstructionManagementService
}

func NewConstructionManagementHandlers(service *service.ConstructionManagementService) *ConstructionManagementHandlers {
	return &ConstructionManagementHandlers{constructionManagementService: service}
}

func (h *ConstructionManagementHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionManagementHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionManagementHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionManagementHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionManagementHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
