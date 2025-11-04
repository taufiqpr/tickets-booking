package handler

import (
	"context"
	pb "ticket-booking/proto/schedule"
	"ticket-booking/schedule-service/internal/service"
)

type GrpcServer struct {
	pb.UnimplementedScheduleServiceServer
	scheduleService service.ScheduleService
}

func NewGrpcServer(scheduleService service.ScheduleService) *GrpcServer {
	return &GrpcServer{scheduleService: scheduleService}
}

func (s *GrpcServer) CreateSchedule(ctx context.Context, req *pb.CreateScheduleRequest) (*pb.CreateScheduleResponse, error) {
	schedule, err := s.scheduleService.CreateSchedule(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateScheduleResponse{Schedule: schedule}, nil
}

func (s *GrpcServer) GetSchedule(ctx context.Context, req *pb.GetScheduleRequest) (*pb.GetScheduleResponse, error) {
	schedule, err := s.scheduleService.GetSchedule(ctx, req.ScheduleId)
	if err != nil {
		return nil, err
	}

	return &pb.GetScheduleResponse{Schedule: schedule}, nil
}

func (s *GrpcServer) ListSchedules(ctx context.Context, req *pb.ListSchedulesRequest) (*pb.ListSchedulesResponse, error) {
	schedules, total, err := s.scheduleService.ListSchedules(ctx, req.Origin, req.Destination, req.DepartureDate, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	return &pb.ListSchedulesResponse{
		Schedules: schedules,
		Total:     total,
	}, nil
}
