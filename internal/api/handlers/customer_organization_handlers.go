package handlers

import (
	"construction-organization-system/internal/service"
	"net/http"
)

type CustomerOrganizationHandlers struct {
	constructionOrganizationService *service.CustomerOrganizationService
}

func NewCustomerOrganizationHandlers(service *service.CustomerOrganizationService) *CustomerOrganizationHandlers {
	return &CustomerOrganizationHandlers{constructionOrganizationService: service}
}

func (h *CustomerOrganizationHandlers) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *CustomerOrganizationHandlers) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *CustomerOrganizationHandlers) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *CustomerOrganizationHandlers) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *CustomerOrganizationHandlers) GetList(w http.ResponseWriter, r *http.Request) {

}
