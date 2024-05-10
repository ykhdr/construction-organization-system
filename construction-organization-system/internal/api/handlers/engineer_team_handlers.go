package handlers

import (
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EngineerTeamHandlers struct {
	engineerTeamService *service.EngineerTeamService
}

func NewEngineerTeamHandlers(service *service.EngineerTeamService) *EngineerTeamHandlers {
	return &EngineerTeamHandlers{engineerTeamService: service}
}

func (h *EngineerTeamHandlers) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get engineer team id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	team, err := h.engineerTeamService.Get(id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting engineer team")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *EngineerTeamHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerTeamHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerTeamHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *EngineerTeamHandlers) GetList(w http.ResponseWriter, r *http.Request) {

	teams, err := h.engineerTeamService.GetList()
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting engineer team list")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(teams)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error encoding engineer team list")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
