package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type ConstructionProjectHandlers struct {
	constructionProjectService *service.ConstructionProjectService
}

func NewConstructionProjectHandlers(service *service.ConstructionProjectService) *ConstructionProjectHandlers {
	return &ConstructionProjectHandlers{constructionProjectService: service}
}

func (h *ConstructionProjectHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionProjectHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
