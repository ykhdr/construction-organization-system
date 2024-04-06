package api

import (
	"construction-organization-system/internal/api/handlers"
	"construction-organization-system/internal/database/repository/postgres"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/service"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	listenAddr                     string
	router                         *mux.Router
	buildingOrganizationHandlers   *handlers.BuildingOrganizationHandlers
	buildingSiteHandlers           *handlers.BuildingSiteHandlers
	constructionContractHandlers   *handlers.ConstructionContractHandlers
	constructionMachineryHandlers  *handlers.ConstructionMachineryHandlers
	constructionManagementHandlers *handlers.ConstructionManagementHandlers
	constructionProjectHandlers    *handlers.ConstructionProjectHandlers
	constructionTeamHandlers       *handlers.ConstructionTeamHandlers
	constructionWorkerHandlers     *handlers.ConstructionWorkerHandlers
	customerOrganizationHandlers   *handlers.CustomerOrganizationHandlers
	employeeHandlers               *handlers.EmployeeHandlers
	engineerTeamHandlers           *handlers.EngineerTeamHandlers
	engineerWorkerHandlers         *handlers.EngineerWorkerHandlers
	estimateHandlers               *handlers.EstimateHandlers
	reportHandlers                 *handlers.ReportHandlers
	workScheduleHandlers           *handlers.WorkScheduleHandlers
}

func NewServer(listenAddr string, db *sql.DB) *Server {
	buildingOrganizationService := service.NewBuildingOrganizationService(postgres.NewBuildingOrganizationRepository(db))
	buildingSiteService := service.NewBuildingSiteService(postgres.NewBuildingSiteRepository(db))
	constructionContractService := service.NewConstructionContractService(postgres.NewConstructionContractRepository(db))
	constructionMachineryService := service.NewConstructionMachineryService(postgres.NewConstructionMachineryRepository(db))
	constructionManagementService := service.NewConstructionManagementService(postgres.NewConstructionManagementRepository(db))
	constructionProjectService := service.NewConstructionProjectService(postgres.NewConstructionProjectRepository(db))
	constructionTeamService := service.NewConstructionTeamService(postgres.NewConstructionTeamRepository(db))
	constructionWorkerService := service.NewConstructionWorkerService(postgres.NewConstructionWorkerRepository(db))
	customerOrganizationService := service.NewCustomerOrganizationService(postgres.NewCustomerOrganizationRepository(db))
	employeeService := service.NewEmployeeService(postgres.NewEmployeeRepository(db))
	engineerTeamService := service.NewEngineerTeamService(postgres.NewEngineerTeamRepository(db))
	engineerWorkerService := service.NewEngineerWorkerService(postgres.NewEngineerWorkerRepository(db))
	estimateService := service.NewEstimateService(postgres.NewEstimateRepository(db))
	reportService := service.NewReportService(postgres.NewReportRepository(db))
	workScheduleService := service.NewWorkScheduleService(postgres.NewWorkScheduleRepository(db))

	buildingOrganizationHandlers := handlers.NewBuildingOrganizationHandlers(buildingOrganizationService)
	buildingSiteHandlers := handlers.NewBuildingSiteHandlers(buildingSiteService, engineerWorkerService)
	constructionContractHandlers := handlers.NewConstructionContractHandlers(constructionContractService)
	constructionMachineryHandlers := handlers.NewConstructionMachineryHandlers(constructionMachineryService)
	constructionManagementHandlers := handlers.NewConstructionManagementHandlers(constructionManagementService, engineerWorkerService)
	constructionProjectHandlers := handlers.NewConstructionProjectHandlers(constructionProjectService)
	constructionTeamHandlers := handlers.NewConstructionTeamHandlers(constructionTeamService)
	constructionWorkerHandlers := handlers.NewConstructionWorkerHandlers(constructionWorkerService)
	customerOrganizationHandlers := handlers.NewCustomerOrganizationHandlers(customerOrganizationService)
	employeeHandlers := handlers.NewEmployeeHandlers(employeeService)
	engineerTeamHandlers := handlers.NewEngineerTeamHandlers(engineerTeamService)
	engineerWorkerHandlers := handlers.NewEngineerWorkerHandlers(engineerWorkerService)
	estimateHandlers := handlers.NewEstimateHandlers(estimateService)
	reportHandlers := handlers.NewReportHandlers(reportService)
	workScheduleHandlers := handlers.NewWorkScheduleHandlers(workScheduleService)

	return &Server{
		listenAddr:                     listenAddr,
		router:                         &mux.Router{},
		buildingOrganizationHandlers:   buildingOrganizationHandlers,
		buildingSiteHandlers:           buildingSiteHandlers,
		constructionContractHandlers:   constructionContractHandlers,
		constructionMachineryHandlers:  constructionMachineryHandlers,
		constructionManagementHandlers: constructionManagementHandlers,
		constructionProjectHandlers:    constructionProjectHandlers,
		constructionTeamHandlers:       constructionTeamHandlers,
		constructionWorkerHandlers:     constructionWorkerHandlers,
		customerOrganizationHandlers:   customerOrganizationHandlers,
		employeeHandlers:               employeeHandlers,
		engineerTeamHandlers:           engineerTeamHandlers,
		engineerWorkerHandlers:         engineerWorkerHandlers,
		estimateHandlers:               estimateHandlers,
		reportHandlers:                 reportHandlers,
		workScheduleHandlers:           workScheduleHandlers,
	}

}

func (s *Server) InitializeRoutes() {
	api := s.router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/building_organization", s.buildingOrganizationHandlers.GetList).Methods("GET")
	api.HandleFunc("/building_organization/{id:[0-9]+}", s.buildingOrganizationHandlers.Get).Methods("GET")
	api.HandleFunc("/building_organization", s.buildingOrganizationHandlers.Create).Methods("POST")
	api.HandleFunc("/building_organization/{id:[0-9]+}", s.buildingOrganizationHandlers.Update).Methods("PUT")
	api.HandleFunc("/building_organization/{id:[0-9]+}", s.buildingOrganizationHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/building_site", s.buildingSiteHandlers.GetList).Methods("GET")
	api.HandleFunc("/building_site/{site-id:[0-9]+}", s.buildingSiteHandlers.Get).Methods("GET")
	api.HandleFunc("/building_site", s.buildingSiteHandlers.Create).Methods("POST")
	api.HandleFunc("/building_site/{site-id:[0-9]+}", s.buildingSiteHandlers.Update).Methods("PUT")
	api.HandleFunc("/building_site/{site-id:[0-9]+}", s.buildingSiteHandlers.Delete).Methods("DELETE")
	api.HandleFunc("/building_site/{site-id:[0-9]+}/manager", s.buildingSiteHandlers.GetManager).Methods("GET")

	api.HandleFunc("/construction_contract", s.constructionContractHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_contract/{id:[0-9]+}", s.constructionContractHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_contract", s.constructionContractHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_contract/{id:[0-9]+}", s.constructionContractHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_contract/{id:[0-9]+}", s.constructionContractHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/construction_machinery", s.constructionMachineryHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_machinery/{id:[0-9]+}", s.constructionMachineryHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_machinery", s.constructionMachineryHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_machinery/{id:[0-9]+}", s.constructionMachineryHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_machinery/{id:[0-9]+}", s.constructionMachineryHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/construction_management", s.constructionManagementHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_management/{management-id:[0-9]+}", s.constructionManagementHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_management", s.constructionManagementHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_management/{management-id:[0-9]+}", s.constructionManagementHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_management/{management-id:[0-9]+}", s.constructionManagementHandlers.Delete).Methods("DELETE")
	api.HandleFunc("/construction_management/{management-id:[0-9]+}/manager", s.constructionManagementHandlers.GetManager).Methods("GET")

	api.HandleFunc("/construction_project", s.constructionProjectHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_project/{id:[0-9]+}", s.constructionProjectHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_project", s.constructionProjectHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_project/{id:[0-9]+}", s.constructionProjectHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_project/{id:[0-9]+}", s.constructionProjectHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/construction_team", s.constructionTeamHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_team/{id:[0-9]+}", s.constructionTeamHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_team", s.constructionTeamHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_team/{id:[0-9]+}", s.constructionTeamHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_team/{id:[0-9]+}", s.constructionTeamHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/construction_worker", s.constructionWorkerHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_worker/{id:[0-9]+}", s.constructionWorkerHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_worker", s.constructionWorkerHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_worker/{id:[0-9]+}", s.constructionWorkerHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_worker/{id:[0-9]+}", s.constructionWorkerHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/customer_organization", s.customerOrganizationHandlers.GetList).Methods("GET")
	api.HandleFunc("/customer_organization/{id:[0-9]+}", s.customerOrganizationHandlers.Get).Methods("GET")
	api.HandleFunc("/customer_organization", s.customerOrganizationHandlers.Create).Methods("POST")
	api.HandleFunc("/customer_organization/{id:[0-9]+}", s.customerOrganizationHandlers.Update).Methods("PUT")
	api.HandleFunc("/customer_organization/{id:[0-9]+}", s.customerOrganizationHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/employee", s.employeeHandlers.GetList).Methods("GET")
	api.HandleFunc("/employee/{id:[0-9]+}", s.employeeHandlers.Get).Methods("GET")
	api.HandleFunc("/employee", s.employeeHandlers.Create).Methods("POST")
	api.HandleFunc("/employee/{id:[0-9]+}", s.employeeHandlers.Update).Methods("PUT")
	api.HandleFunc("/employee/{id:[0-9]+}", s.employeeHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/engineer_team", s.engineerTeamHandlers.GetList).Methods("GET")
	api.HandleFunc("/engineer_team/{id:[0-9]+}", s.engineerTeamHandlers.Get).Methods("GET")
	api.HandleFunc("/engineer_team", s.engineerTeamHandlers.Create).Methods("POST")
	api.HandleFunc("/engineer_team/{id:[0-9]+}", s.engineerTeamHandlers.Update).Methods("PUT")
	api.HandleFunc("/engineer_team/{id:[0-9]+}", s.engineerTeamHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/engineer_worker", s.engineerWorkerHandlers.GetList).Methods("GET")
	api.HandleFunc("/engineer_worker/{id:[0-9]+}", s.engineerWorkerHandlers.Get).Methods("GET")
	api.HandleFunc("/engineer_worker", s.engineerWorkerHandlers.Create).Methods("POST")
	api.HandleFunc("/engineer_worker/{id:[0-9]+}", s.engineerWorkerHandlers.Update).Methods("PUT")
	api.HandleFunc("/engineer_worker/{id:[0-9]+}", s.engineerWorkerHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/estimate", s.estimateHandlers.GetList).Methods("GET")
	api.HandleFunc("/estimate/{id:[0-9]+}", s.estimateHandlers.Get).Methods("GET")
	api.HandleFunc("/estimate", s.estimateHandlers.Create).Methods("POST")
	api.HandleFunc("/estimate/{id:[0-9]+}", s.estimateHandlers.Update).Methods("PUT")
	api.HandleFunc("/estimate/{id:[0-9]+}", s.estimateHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/report", s.reportHandlers.GetList).Methods("GET")
	api.HandleFunc("/report/{id:[0-9]+}", s.reportHandlers.Get).Methods("GET")
	api.HandleFunc("/report", s.reportHandlers.Create).Methods("POST")
	api.HandleFunc("/report/{id:[0-9]+}", s.reportHandlers.Update).Methods("PUT")
	api.HandleFunc("/report/{id:[0-9]+}", s.reportHandlers.Delete).Methods("DELETE")

	api.HandleFunc("/work_schedule", s.workScheduleHandlers.GetList).Methods("GET")
	api.HandleFunc("/work_schedule/{id:[0-9]+}", s.workScheduleHandlers.Get).Methods("GET")
	api.HandleFunc("/work_schedule", s.workScheduleHandlers.Create).Methods("POST")
	api.HandleFunc("/work_schedule/{id:[0-9]+}", s.workScheduleHandlers.Update).Methods("PUT")
	api.HandleFunc("/work_schedule/{id:[0-9]+}", s.workScheduleHandlers.Delete).Methods("DELETE")
}

func (s *Server) Start() {
	log.Logger.Fatal(http.ListenAndServe(s.listenAddr, s.router))
}
