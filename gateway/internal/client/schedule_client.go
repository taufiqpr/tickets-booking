package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "ticket-booking/proto/schedule"
)

type ScheduleClient struct {
	client pb.ScheduleServiceClient
}

func NewScheduleClient(host string, port int) (*ScheduleClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewScheduleServiceClient(conn)
	return &ScheduleClient{client: client}, nil
}

func (c *ScheduleClient) GetSchedule(ctx context.Context, scheduleID int64) (*pb.GetScheduleResponse, error) {
	return c.client.GetSchedule(ctx, &pb.GetScheduleRequest{
		ScheduleId: scheduleID,
	})
}

func (c *ScheduleClient) ListSchedules(ctx context.Context, origin, destination, departureDate string, page, limit int32) (*pb.ListSchedulesResponse, error) {
	return c.client.ListSchedules(ctx, &pb.ListSchedulesRequest{
		Origin:        origin,
		Destination:   destination,
		DepartureDate: departureDate,
		Page:          page,
		Limit:         limit,
	})
}

func (c *ScheduleClient) CreateSchedule(ctx context.Context, trainID int64, origin, destination, departureTime, arrivalTime string, price float64) (*pb.CreateScheduleResponse, error) {
	return c.client.CreateSchedule(ctx, &pb.CreateScheduleRequest{
		TrainId:       trainID,
		Origin:        origin,
		Destination:   destination,
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		Price:         price,
	})
}
