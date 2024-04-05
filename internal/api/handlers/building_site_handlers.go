package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BuildingSiteHandlers struct {
	buildingSiteService *service.BuildingSiteService
}

func NewBuildingSiteHandlers(service *service.BuildingSiteService) *BuildingSiteHandlers {
	return &BuildingSiteHandlers{buildingSiteService: service}
}

func (h *BuildingSiteHandlers) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buildingSite, err := h.buildingSiteService.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(buildingSite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *BuildingSiteHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *BuildingSiteHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
