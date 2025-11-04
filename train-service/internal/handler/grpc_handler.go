package handler

import (
	"context"
	pb "ticket-booking/proto/train"
	"ticket-booking/train-service/internal/service"
)

type GrpcServer struct {
	pb.UnimplementedTrainServiceServer
	trainService service.TrainService
}

func NewGrpcServer(trainService service.TrainService) *GrpcServer {
	return &GrpcServer{trainService: trainService}
}

func (s *GrpcServer) CreateTrain(ctx context.Context, req *pb.CreateTrainRequest) (*pb.CreateTrainResponse, error) {
	train, err := s.trainService.CreateTrain(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTrainResponse{Train: train}, nil
}

func (s *GrpcServer) GetTrain(ctx context.Context, req *pb.GetTrainRequest) (*pb.GetTrainResponse, error) {
	train, err := s.trainService.GetTrain(ctx, req.TrainId)
	if err != nil {
		return nil, err
	}

	return &pb.GetTrainResponse{Train: train}, nil
}

func (s *GrpcServer) ListTrains(ctx context.Context, req *pb.ListTrainsRequest) (*pb.ListTrainsResponse, error) {
	trains, total, err := s.trainService.ListTrains(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	return &pb.ListTrainsResponse{
		Trains: trains,
		Total:  total,
	}, nil
}

func (s *GrpcServer) UpdateTrain(ctx context.Context, req *pb.UpdateTrainRequest) (*pb.UpdateTrainResponse, error) {
	train, err := s.trainService.UpdateTrain(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTrainResponse{Train: train}, nil
}

func (s *GrpcServer) DeleteTrain(ctx context.Context, req *pb.DeleteTrainRequest) (*pb.DeleteTrainResponse, error) {
	err := s.trainService.DeleteTrain(ctx, req.TrainId)
	if err != nil {
		return &pb.DeleteTrainResponse{Success: false}, err
	}

	return &pb.DeleteTrainResponse{Success: true}, nil
}
