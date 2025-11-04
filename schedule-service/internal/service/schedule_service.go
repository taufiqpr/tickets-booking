package service

import (
	"context"
	pb "ticket-booking/proto/schedule"
	"ticket-booking/schedule-service/internal/repository"
)

type ScheduleService interface {
	CreateSchedule(ctx context.Context, req *pb.CreateScheduleRequest) (*pb.Schedule, error)
	GetSchedule(ctx context.Context, id int64) (*pb.Schedule, error)
	ListSchedules(ctx context.Context, origin, destination, departureDate string, page, limit int32) ([]*pb.Schedule, int32, error)
}

type scheduleService struct {
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleService(scheduleRepo repository.ScheduleRepository) ScheduleService {
	return &scheduleService{scheduleRepo: scheduleRepo}
}

func (s *scheduleService) CreateSchedule(ctx context.Context, req *pb.CreateScheduleRequest) (*pb.Schedule, error) {
	return s.scheduleRepo.Create(ctx, req)
}

func (s *scheduleService) GetSchedule(ctx context.Context, id int64) (*pb.Schedule, error) {
	return s.scheduleRepo.GetByID(ctx, id)
}

func (s *scheduleService) ListSchedules(ctx context.Context, origin, destination, departureDate string, page, limit int32) ([]*pb.Schedule, int32, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return s.scheduleRepo.List(ctx, origin, destination, departureDate, page, limit)
}
