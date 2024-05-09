package service

import (
	"construction-organization-system/internal/database/repository"
	"construction-organization-system/internal/model"
	"context"
	"encoding/json"
	"io"
)

type WorkScheduleService struct {
	workScheduleRepository repository.WorkScheduleRepository
}

func NewWorkScheduleService(repo repository.WorkScheduleRepository) *WorkScheduleService {
	return &WorkScheduleService{workScheduleRepository: repo}
}

func (s *WorkScheduleService) GetList() ([]*model.WorkSchedule, error) {
	schedules, err := s.workScheduleRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (s *WorkScheduleService) Create(body io.Reader) error {
	var workSchedule model.WorkSchedule

	err := json.NewDecoder(body).Decode(&workSchedule)
	if err != nil {
		return err
	}

	id, err := s.workScheduleRepository.Save(context.Background(), workSchedule)
	if err != nil {
		return err
	}
	workSchedule.ID = id
	return nil
}

func (s *WorkScheduleService) Get(id int) (*model.WorkSchedule, error) {
	return s.workScheduleRepository.Find(context.Background(), id)
}

func (s *WorkScheduleService) Delete(id int) error {
	return s.workScheduleRepository.Delete(context.Background(), id)
}

func (s *WorkScheduleService) Update(id int, body io.ReadCloser) error {
	var workSchedule model.WorkSchedule

	err := json.NewDecoder(body).Decode(&workSchedule)
	if err != nil {
		return err
	}

	return s.workScheduleRepository.Update(context.Background(), workSchedule)
}
