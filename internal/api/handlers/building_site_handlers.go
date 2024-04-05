package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type BuildingSiteHandlers struct {
	buildingSiteService *service.BuildingSiteService
}

func NewBuildingSiteHandlers(service *service.BuildingSiteService) *BuildingSiteHandlers {
	return &BuildingSiteHandlers{buildingSiteService: service}
}

func (h *BuildingSiteHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
