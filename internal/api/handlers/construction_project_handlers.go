package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
