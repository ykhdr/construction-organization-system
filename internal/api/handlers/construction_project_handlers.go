package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
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

func (h *ConstructionProjectHandlers) GetWorkSchedules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	workSchedules, err := h.constructionProjectService.GetWorkSchedules(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workSchedules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetConstructionTeams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	constructionTeams, err := h.constructionProjectService.GetConstructionTeams(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(constructionTeams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetMachines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
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

	machines, err := h.constructionProjectService.GetMachines(projectId, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(machines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetEstimate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	estimate, err := h.constructionProjectService.GetEstimate(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(estimate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionProjectHandlers) GetReports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reports, err := h.constructionProjectService.GetReports(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reports)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
