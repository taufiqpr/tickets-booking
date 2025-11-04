package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "ticket-booking/proto/train"
)

type TrainClient struct {
	client pb.TrainServiceClient
}

func NewTrainClient(host string, port int) (*TrainClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewTrainServiceClient(conn)
	return &TrainClient{client: client}, nil
}

func (c *TrainClient) GetTrain(ctx context.Context, trainID int64) (*pb.GetTrainResponse, error) {
	return c.client.GetTrain(ctx, &pb.GetTrainRequest{
		TrainId: trainID,
	})
}

func (c *TrainClient) ListTrains(ctx context.Context, page, limit int32) (*pb.ListTrainsResponse, error) {
	return c.client.ListTrains(ctx, &pb.ListTrainsRequest{
		Page:  page,
		Limit: limit,
	})
}

func (c *TrainClient) CreateTrain(ctx context.Context, name, trainType string, capacity int32) (*pb.CreateTrainResponse, error) {
	return c.client.CreateTrain(ctx, &pb.CreateTrainRequest{
		Name:     name,
		Type:     trainType,
		Capacity: capacity,
	})
}

func (c *TrainClient) UpdateTrain(ctx context.Context, trainID int64, name, trainType, status string, capacity int32) (*pb.UpdateTrainResponse, error) {
	return c.client.UpdateTrain(ctx, &pb.UpdateTrainRequest{
		TrainId:  trainID,
		Name:     name,
		Type:     trainType,
		Capacity: capacity,
		Status:   status,
	})
}

func (c *TrainClient) DeleteTrain(ctx context.Context, trainID int64) (*pb.DeleteTrainResponse, error) {
	return c.client.DeleteTrain(ctx, &pb.DeleteTrainRequest{
		TrainId: trainID,
	})
}
