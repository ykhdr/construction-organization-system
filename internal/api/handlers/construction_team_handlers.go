package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
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

func (h *ConstructionTeamHandlers) GetWorkTypes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if _, err := time.Parse("2006-01-02", startDate); startDate != "" && err != nil {
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", endDate); endDate != "" && err != nil {
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	workTypes, err := h.constructionTeamService.GetWorkTypes(teamId, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workTypes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
