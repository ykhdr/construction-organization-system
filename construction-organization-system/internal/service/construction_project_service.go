package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
)

type ConstructionProjectService struct {
	constructionProjectRepository   repository.ConstructionProjectRepository
	workScheduleRepository          repository.WorkScheduleRepository
	constructionTeamRepository      repository.ConstructionTeamRepository
	constructionMachineryRepository repository.ConstructionMachineryRepository
	estimateRepository              repository.EstimateRepository
	reportRepository                repository.ReportRepository
	workTypeRepository              repository.WorkTypeRepository
	schoolRepository                repository.SchoolRepository
	apartmentHouseRepository        repository.ApartmentHouseRepository
	bridgeRepository                repository.BridgeRepository
}

func NewConstructionProjectService(projectRepository repository.ConstructionProjectRepository,
	workScheduleRepository repository.WorkScheduleRepository,
	constructionTeamRepository repository.ConstructionTeamRepository,
	constructionMachineryRepository repository.ConstructionMachineryRepository,
	estimateRepository repository.EstimateRepository,
	reportRepository repository.ReportRepository,
	workTypeRepository repository.WorkTypeRepository,
	schoolRepository repository.SchoolRepository,
	apartmentHouseRepository repository.ApartmentHouseRepository,
	bridgeRepository repository.BridgeRepository) *ConstructionProjectService {
	return &ConstructionProjectService{projectRepository, workScheduleRepository, constructionTeamRepository,
		constructionMachineryRepository, estimateRepository, reportRepository, workTypeRepository, schoolRepository,
		apartmentHouseRepository, bridgeRepository}
}

func (s *ConstructionProjectService) GetWorkSchedules(projectID int) ([]*model.WorkSchedule, error) {
	workSchedules, err := s.workScheduleRepository.FindByProject(context.Background(), projectID)
	if err != nil {
		return nil, err
	}
	return workSchedules, nil
}

func (s *ConstructionProjectService) GetConstructionTeams(projectID int) ([]*model.ConstructionTeam, error) {
	teams, err := s.constructionTeamRepository.FindByProject(context.Background(), projectID)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *ConstructionProjectService) GetMachines(projectID int, startDate, endDate string) ([]*model.ConstructionMachinery, error) {
	var machines []*model.ConstructionMachinery
	var err error

	if startDate == "" && endDate == "" {
		machines, err = s.constructionMachineryRepository.GetByProject(context.Background(), projectID)
	} else {
		machines, err = s.constructionMachineryRepository.GetByProjectWithPeriod(context.Background(), projectID, startDate, endDate)
	}

	if err != nil {
		return nil, err
	}

	return machines, nil
}

func (s *ConstructionProjectService) GetEstimate(projectID int) (*model.Estimate, error) {
	estimate, err := s.estimateRepository.Find(context.Background(), projectID)

	if err != nil {
		return nil, err
	}

	return estimate, nil
}

func (s *ConstructionProjectService) GetReports(projectId int) ([]*model.Report, error) {
	reports, err := s.reportRepository.FindByProjectID(context.Background(), projectId)

	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (s *ConstructionProjectService) GetProjects(workTypeID int, startDate, endDate string) ([]*model.ConstructionProject, error) {
	var projects []*model.ConstructionProject
	var err error
	ctx := context.Background()

	if workTypeID != 0 {
		projects, err = s.constructionProjectRepository.FindByWorkTypeWithPeriod(ctx, workTypeID, startDate, endDate)
	} else {
		projects, err = s.constructionProjectRepository.FindAll(context.Background())
	}

	if err != nil {
		return nil, err
	}

	return projects, err
}

func (s *ConstructionProjectService) GetExceededDeadlinesWorks(projectId int) ([]*model.WorkType, error) {
	workTypes, err := s.workTypeRepository.FindByProjectWithExceededWorkDeadlines(context.Background(), projectId)
	if err != nil {
		return nil, err
	}

	return workTypes, nil
}

func (s *ConstructionProjectService) GetSchoolList() ([]*model.School, error) {
	schools, err := s.schoolRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}

	return schools, nil
}

func (s *ConstructionProjectService) CreateSchool(school model.School) error {
	_, err := s.schoolRepository.Save(context.Background(), school)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) GetSchool(projectId int) (*model.School, error) {
	school, err := s.schoolRepository.Find(context.Background(), projectId)
	if err != nil {
		return nil, err
	}

	return school, nil
}

func (s *ConstructionProjectService) UpdateSchool(school model.School) error {
	err := s.schoolRepository.Update(context.Background(), school)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) DeleteSchool(projectId int) error {
	err := s.schoolRepository.Delete(context.Background(), projectId)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) GetApartmentHouseList() ([]*model.ApartmentHouse, error) {
	houses, err := s.apartmentHouseRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}

	return houses, nil
}

func (s *ConstructionProjectService) GetApartmentHouse(projectId int) (*model.ApartmentHouse, error) {
	house, err := s.apartmentHouseRepository.Find(context.Background(), projectId)
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (s *ConstructionProjectService) UpdateApartmentHouse(house model.ApartmentHouse) error {
	err := s.apartmentHouseRepository.Update(context.Background(), house)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) DeleteApartmentHouse(projectId int) error {
	err := s.apartmentHouseRepository.Delete(context.Background(), projectId)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) CreateApartmentHouse(house model.ApartmentHouse) error {
	_, err := s.apartmentHouseRepository.Save(context.Background(), house)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) GetBridgeList() ([]*model.Bridge, error) {
	bridges, err := s.bridgeRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}

	return bridges, nil
}

func (s *ConstructionProjectService) GetBridge(projectId int) (*model.Bridge, error) {
	bridge, err := s.bridgeRepository.Find(context.Background(), projectId)
	if err != nil {
		return nil, err
	}

	return bridge, nil
}

func (s *ConstructionProjectService) CreateBridge(bridge model.Bridge) error {
	_, err := s.bridgeRepository.Save(context.Background(), bridge)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) UpdateBridge(bridge model.Bridge) error {
	err := s.bridgeRepository.Update(context.Background(), bridge)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) DeleteBridge(projectId int) error {
	err := s.bridgeRepository.Delete(context.Background(), projectId)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConstructionProjectService) GetProject(projectId int) (*model.ConstructionProject, error) {
	project, err := s.constructionProjectRepository.Find(context.Background(), projectId)
	if err != nil {
		return nil, err
	}

	return project, nil
}
