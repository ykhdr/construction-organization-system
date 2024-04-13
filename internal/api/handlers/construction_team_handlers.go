package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (h *ConstructionTeamHandlers) GetWorkers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	workers, err := h.constructionTeamService.GetWorkers(teamId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
