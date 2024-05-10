package handlers

import (
	"construction-organization-system/internal/log"
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
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get team id")
		http.Error(w, "Invalid id format. Use integer value.", http.StatusBadRequest)
		return
	}

	team, err := h.constructionTeamService.Get(id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting team")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(team)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error encoding team")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionTeamHandlers) Create(w http.ResponseWriter, r *http.Request) {
	err := h.constructionTeamService.Create(r.Body)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error creating team")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (h *ConstructionTeamHandlers) Update(w http.ResponseWriter, r *http.Request) {
	err := h.constructionTeamService.Update(r.Body)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error updating team")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ConstructionTeamHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get team id")
		http.Error(w, "Invalid id format. Use integer value.", http.StatusBadRequest)
		return
	}

	err = h.constructionTeamService.Delete(id)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error deleting team")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ConstructionTeamHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	var workTypeID int
	var err error

	workType := r.URL.Query().Get("work_type")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if workType == "" {
		workTypeID = 0
	} else if workTypeID, err = strconv.Atoi(workType); err != nil {
		log.Logger.WithError(err).Errorln("Error get work type")
		http.Error(w, "Invalid work type format. Use integer value.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", startDate); startDate != "" && err != nil {
		log.Logger.WithError(err).Errorln("Error get start date")
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", endDate); endDate != "" && err != nil {
		log.Logger.WithError(err).Errorln("Error get end date")
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	teams, err := h.constructionTeamService.GetTeams(workTypeID, startDate, endDate)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting teams")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(teams)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error encoding teams")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *ConstructionTeamHandlers) GetWorkers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get team id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	workers, err := h.constructionTeamService.GetWorkers(teamId)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting workers")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workers)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error encoding workers")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ConstructionTeamHandlers) GetWorkTypes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Errorln("Error get team id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if _, err := time.Parse("2006-01-02", startDate); startDate != "" && err != nil {
		log.Logger.WithError(err).Errorln("Error get start date")
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", endDate); endDate != "" && err != nil {
		log.Logger.WithError(err).Errorln("Error get end date")
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	workTypes, err := h.constructionTeamService.GetWorkTypes(teamId, startDate, endDate)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error getting work types")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(workTypes)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error encoding work types")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
