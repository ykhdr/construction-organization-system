package api

import (
	"construction-organization-system/internal/api/handlers"
	"construction-organization-system/internal/database/repository/postgres"
	"construction-organization-system/internal/log"
	"construction-organization-system/internal/net"
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
	workTypeHandlers               *handlers.WorkTypeHandlers
}

func NewServer(listenAddr string, db *sql.DB) *Server {
	buildingOrganizationRepository := postgres.NewBuildingOrganizationRepository(db)
	buildingSiteRepository := postgres.NewBuildingSiteRepository(db)
	constructionContractRepository := postgres.NewConstructionContractRepository(db)
	constructionMachineryRepository := postgres.NewConstructionMachineryRepository(db)
	constructionManagementRepository := postgres.NewConstructionManagementRepository(db)
	constructionProjectRepository := postgres.NewConstructionProjectRepository(db)
	constructionTeamRepository := postgres.NewConstructionTeamRepository(db)
	constructionWorkerRepository := postgres.NewConstructionWorkerRepository(db)
	customerOrganizationRepository := postgres.NewCustomerOrganizationRepository(db)
	employeeRepository := postgres.NewEmployeeRepository(db)
	engineerTeamRepository := postgres.NewEngineerTeamRepository(db)
	engineerWorkerRepository := postgres.NewEngineerWorkerRepository(db)
	estimateRepository := postgres.NewEstimateRepository(db)
	reportRepository := postgres.NewReportRepository(db)
	workScheduleRepository := postgres.NewWorkScheduleRepository(db)
	workTypeRepository := postgres.NewWorkTypeRepository(db)
	schoolRepository := postgres.NewSchoolRepository(db)
	apartmentHouseRepository := postgres.NewApartmentHouseRepository(db)
	bridgeRepository := postgres.NewBridgeRepository(db)
	materialRepository := postgres.NewMaterialRepository(db)

	buildingOrganizationService := service.NewBuildingOrganizationService(buildingOrganizationRepository)
	buildingSiteService := service.NewBuildingSiteService(
		buildingSiteRepository,
		engineerWorkerRepository,
		constructionProjectRepository,
	)
	constructionContractService := service.NewConstructionContractService(constructionContractRepository)
	constructionMachineryService := service.NewConstructionMachineryService(constructionMachineryRepository)
	constructionManagementService := service.NewConstructionManagementService(
		constructionManagementRepository,
		engineerWorkerRepository,
		constructionProjectRepository,
		constructionMachineryRepository,
	)
	constructionProjectService := service.NewConstructionProjectService(
		constructionProjectRepository,
		workScheduleRepository,
		constructionTeamRepository,
		constructionMachineryRepository,
		estimateRepository,
		reportRepository,
		workTypeRepository,
		schoolRepository,
		apartmentHouseRepository,
		bridgeRepository,
	)
	constructionTeamService := service.NewConstructionTeamService(
		constructionTeamRepository,
		constructionWorkerRepository,
		workTypeRepository,
	)
	constructionWorkerService := service.NewConstructionWorkerService(constructionWorkerRepository)
	customerOrganizationService := service.NewCustomerOrganizationService(customerOrganizationRepository)
	employeeService := service.NewEmployeeService(employeeRepository)
	engineerTeamService := service.NewEngineerTeamService(engineerTeamRepository)
	engineerWorkerService := service.NewEngineerWorkerService(engineerWorkerRepository)
	estimateService := service.NewEstimateService(estimateRepository, materialRepository)
	reportService := service.NewReportService(reportRepository)
	workScheduleService := service.NewWorkScheduleService(workScheduleRepository)
	workTypeService := service.NewWorkTypeService(workTypeRepository)

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
	workTypeHandlers := handlers.NewWorkTypeHandlers(workTypeService)

	router := mux.NewRouter()
	router.Use(net.LoggerMiddleware)

	return &Server{
		listenAddr:                     listenAddr,
		router:                         router,
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
		workTypeHandlers:               workTypeHandlers,
	}

}

func (s *Server) InitializeRoutes() {
	api := s.router.PathPrefix("/api/v1").Subrouter()

	log.Logger.Infoln("Initializing routes...")

	api.HandleFunc("/building_organization", s.buildingOrganizationHandlers.GetList).Methods("GET")
	api.HandleFunc("/building_organization/{id:[0-9]+}", s.buildingOrganizationHandlers.Get).Methods("GET")
	api.HandleFunc("/building_organization", s.buildingOrganizationHandlers.Create).Methods("POST")
	api.HandleFunc("/building_organization/{id:[0-9]+}", s.buildingOrganizationHandlers.Update).Methods("PUT")
	api.HandleFunc("/building_organization/{id:[0-9]+}", s.buildingOrganizationHandlers.Delete).Methods("DELETE")
	api.HandleFunc("/building_organization/{id:[0-9]+}/projects", s.buildingOrganizationHandlers.GetProjects).Methods("GET")

	api.HandleFunc("/building_site", s.buildingSiteHandlers.GetList).Methods("GET")
	api.HandleFunc("/building_site/{id:[0-9]+}", s.buildingSiteHandlers.Get).Methods("GET")
	api.HandleFunc("/building_site", s.buildingSiteHandlers.Create).Methods("POST")
	api.HandleFunc("/building_site/{id:[0-9]+}", s.buildingSiteHandlers.Update).Methods("PUT")
	api.HandleFunc("/building_site/{id:[0-9]+}", s.buildingSiteHandlers.Delete).Methods("DELETE")
	api.HandleFunc("/building_site/{id:[0-9]+}/manager", s.buildingSiteHandlers.GetManager).Methods("GET")
	api.HandleFunc("/building_site/{id:[0-9]+}/engineers", s.buildingSiteHandlers.GetEngineers).Methods("GET")
	api.HandleFunc("/building_site/{id:[0-9]+}/projects", s.buildingSiteHandlers.GetProjects).Methods("GET")

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
	api.HandleFunc("/construction_management/{id:[0-9]+}", s.constructionManagementHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_management", s.constructionManagementHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_management/{id:[0-9]+}", s.constructionManagementHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_management/{id:[0-9]+}", s.constructionManagementHandlers.Delete).Methods("DELETE")
	api.HandleFunc("/construction_management/{id:[0-9]+}/manager", s.constructionManagementHandlers.GetManager).Methods("GET")
	api.HandleFunc("/construction_management/{id:[0-9]+}/engineers", s.constructionManagementHandlers.GetEngineers).Methods("GET")
	api.HandleFunc("/construction_management/{id:[0-9]+}/projects", s.constructionManagementHandlers.GetProjects).Methods("GET")
	api.HandleFunc("/construction_management/{id:[0-9]+}/machines", s.constructionManagementHandlers.GetMachines).Methods("GET")

	api.HandleFunc("/construction_project", s.constructionProjectHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_project/{id:[0-9]+}", s.constructionProjectHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_project", s.constructionProjectHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_project/{id:[0-9]+}", s.constructionProjectHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_project/{id:[0-9]+}", s.constructionProjectHandlers.Delete).Methods("DELETE")
	api.HandleFunc("/construction_project/{id:[0-9]+}/schedules", s.constructionProjectHandlers.GetWorkSchedules).Methods("GET")
	api.HandleFunc("/construction_project/{id:[0-9]+}/estimate", s.constructionProjectHandlers.GetEstimate).Methods("GET")
	api.HandleFunc("/construction_project/{id:[0-9]+}/construction_teams", s.constructionProjectHandlers.GetConstructionTeams).Methods("GET")
	api.HandleFunc("/construction_project/{id:[0-9]+}/machines", s.constructionProjectHandlers.GetMachines).Methods("GET")
	api.HandleFunc("/construction_project/{id:[0-9]+}/reports", s.constructionProjectHandlers.GetReports).Methods("GET")
	api.HandleFunc("/construction_project/{id:[0-9]+}/exceeded_deadlines_works", s.constructionProjectHandlers.GetExceededDeadlinesWorks).Methods("GET")

	api.HandleFunc("/construction_team", s.constructionTeamHandlers.GetList).Methods("GET")
	api.HandleFunc("/construction_team/{id:[0-9]+}", s.constructionTeamHandlers.Get).Methods("GET")
	api.HandleFunc("/construction_team", s.constructionTeamHandlers.Create).Methods("POST")
	api.HandleFunc("/construction_team/{id:[0-9]+}", s.constructionTeamHandlers.Update).Methods("PUT")
	api.HandleFunc("/construction_team/{id:[0-9]+}", s.constructionTeamHandlers.Delete).Methods("DELETE")
	api.HandleFunc("/construction_team/{id:[0-9]+}", s.constructionTeamHandlers.GetWorkers).Methods("GET")
	api.HandleFunc("/construction_team/{id:[0-9]+}/work_types", s.constructionTeamHandlers.GetWorkTypes).Methods("GET")

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
	api.HandleFunc("/estimate/{id:[0-9]+}/exceeded_usage_material", s.estimateHandlers.GetExceededUsageMaterials).Methods("GET")

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

	api.HandleFunc("/school", s.constructionProjectHandlers.GetSchoolList).Methods("GET")
	api.HandleFunc("/school/{id:[0-9]+}", s.constructionProjectHandlers.GetSchool).Methods("GET")
	api.HandleFunc("/school", s.constructionProjectHandlers.CreateSchool).Methods("POST")
	api.HandleFunc("/school/{id:[0-9]+}", s.constructionProjectHandlers.UpdateSchool).Methods("PUT")
	api.HandleFunc("/school/{id:[0-9]+}", s.constructionProjectHandlers.DeleteSchool).Methods("DELETE")

	api.HandleFunc("/apartment_house", s.constructionProjectHandlers.GetApartmentHouseList).Methods("GET")
	api.HandleFunc("/apartment_house/{id:[0-9]+}", s.constructionProjectHandlers.GetApartmentHouse).Methods("GET")
	api.HandleFunc("/apartment_house", s.constructionProjectHandlers.CreateApartmentHouse).Methods("POST")
	api.HandleFunc("/apartment_house/{id:[0-9]+}", s.constructionProjectHandlers.UpdateApartmentHouse).Methods("PUT")
	api.HandleFunc("/apartment_house/{id:[0-9]+}", s.constructionProjectHandlers.DeleteApartmentHouse).Methods("DELETE")

	api.HandleFunc("/bridge", s.constructionProjectHandlers.GetBridgeList).Methods("GET")
	api.HandleFunc("/bridge/{id:[0-9]+}", s.constructionProjectHandlers.GetBridge).Methods("GET")
	api.HandleFunc("/bridge", s.constructionProjectHandlers.CreateBridge).Methods("POST")
	api.HandleFunc("/bridge/{id:[0-9]+}", s.constructionProjectHandlers.UpdateBridge).Methods("PUT")
	api.HandleFunc("/bridge/{id:[0-9]+}", s.constructionProjectHandlers.DeleteBridge).Methods("DELETE")

	api.HandleFunc("/work_types", s.workTypeHandlers.GetList).Methods("GET")

	log.Logger.Infoln("Routes initialized")
}

func (s *Server) Start() {
	log.Logger.Fatal(http.ListenAndServe(s.listenAddr, s.router))
}
