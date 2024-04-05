package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type ReportHandlers struct {
	reportService *service.ReportService
}

func NewReportHandlers(service *service.ReportService) *ReportHandlers {
	return &ReportHandlers{reportService: service}
}

func (h *ReportHandlers) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *ReportHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ReportHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ReportHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *ReportHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
