package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type ConstructionContractHandlers struct {
	constructionWorkerService *service.ConstructionContractService
}

func NewConstructionContractHandlers(service *service.ConstructionContractService) *ConstructionContractHandlers {
	return &ConstructionContractHandlers{constructionWorkerService: service}
}

func (h *ConstructionContractHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionContractHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionContractHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionContractHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionContractHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
