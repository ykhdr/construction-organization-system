package handlers

import (
	"construction-organization-system/internal/service"
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

}
