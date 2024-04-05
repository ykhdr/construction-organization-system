package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type ConstructionTeamHandlers struct {
	constructionTeamService *service.ConstructionTeamService
}

func NewConstructionTeamHandlers(service *service.ConstructionTeamService) *ConstructionTeamHandlers {
	return &ConstructionTeamHandlers{constructionTeamService: service}
}

func (h *ConstructionTeamHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionTeamHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionTeamHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionTeamHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ConstructionTeamHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
