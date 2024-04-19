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
}

func NewConstructionProjectService(constructionProjectRepo repository.ConstructionProjectRepository, workScheduleRepo repository.WorkScheduleRepository) *ConstructionProjectService {
	return &ConstructionProjectService{constructionProjectRepository: constructionProjectRepo, workScheduleRepository: workScheduleRepo}
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

	if workTypeID == 0 {
		projects, err = s.constructionProjectRepository.FindByWorkTypeWithPeriod(ctx, workTypeID, startDate, endDate)
	} else {
		projects, err = s.constructionProjectRepository.FindAll(context.Background())
	}

	if err != nil {
		return nil, err
	}

	return projects, err
}
