package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type EngineerTeamHandlers struct {
	engineerTeamService *service.EngineerTeamService
}

func NewEngineerTeamHandlers(service *service.EngineerTeamService) *EngineerTeamHandlers {
	return &EngineerTeamHandlers{engineerTeamService: service}
}

func (h *EngineerTeamHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerTeamHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerTeamHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerTeamHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerTeamHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
