package handlers

import (
	"construction-organization-system/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type WorkScheduleHandlers struct {
	workScheduleService *service.WorkScheduleService
}

func NewWorkScheduleHandlers(service *service.WorkScheduleService) *WorkScheduleHandlers {
	return &WorkScheduleHandlers{workScheduleService: service}
}

func (h *WorkScheduleHandlers) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id format. Use integer value.", http.StatusBadRequest)
		return
	}

	schedule, err := h.workScheduleService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *WorkScheduleHandlers) Create(w http.ResponseWriter, r *http.Request) {
	err := h.workScheduleService.Create(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *WorkScheduleHandlers) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id format. Use integer value.", http.StatusBadRequest)
		return
	}
	err = h.workScheduleService.Update(id, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *WorkScheduleHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id format. Use integer value.", http.StatusBadRequest)
		return
	}
	err = h.workScheduleService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *WorkScheduleHandlers) GetList(w http.ResponseWriter, r *http.Request) {
	schedules, err := h.workScheduleService.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(schedules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
