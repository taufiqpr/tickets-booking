package service

import (
	"context"
	pb "ticket-booking/proto/train"
	"ticket-booking/train-service/internal/repository"
)

type TrainService interface {
	CreateTrain(ctx context.Context, req *pb.CreateTrainRequest) (*pb.Train, error)
	GetTrain(ctx context.Context, id int64) (*pb.Train, error)
	ListTrains(ctx context.Context, page, limit int32) ([]*pb.Train, int32, error)
	UpdateTrain(ctx context.Context, req *pb.UpdateTrainRequest) (*pb.Train, error)
	DeleteTrain(ctx context.Context, id int64) error
}

type trainService struct {
	trainRepo repository.TrainRepository
}

func NewTrainService(trainRepo repository.TrainRepository) TrainService {
	return &trainService{trainRepo: trainRepo}
}

func (s *trainService) CreateTrain(ctx context.Context, req *pb.CreateTrainRequest) (*pb.Train, error) {
	return s.trainRepo.Create(ctx, req)
}

func (s *trainService) GetTrain(ctx context.Context, id int64) (*pb.Train, error) {
	return s.trainRepo.GetByID(ctx, id)
}

func (s *trainService) ListTrains(ctx context.Context, page, limit int32) ([]*pb.Train, int32, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return s.trainRepo.List(ctx, page, limit)
}

func (s *trainService) UpdateTrain(ctx context.Context, req *pb.UpdateTrainRequest) (*pb.Train, error) {
	return s.trainRepo.Update(ctx, req)
}

func (s *trainService) DeleteTrain(ctx context.Context, id int64) error {
	return s.trainRepo.Delete(ctx, id)
}
