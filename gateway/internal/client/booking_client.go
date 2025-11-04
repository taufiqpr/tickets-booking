package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "ticket-booking/proto/booking"
)

type BookingClient struct {
	client pb.BookingServiceClient
}

func NewBookingClient(host string, port int) (*BookingClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewBookingServiceClient(conn)
	return &BookingClient{client: client}, nil
}

func (c *BookingClient) CreateBooking(ctx context.Context, userID, scheduleID int64, seatCount int32) (*pb.CreateBookingResponse, error) {
	return c.client.CreateBooking(ctx, &pb.CreateBookingRequest{
		UserId:     userID,
		ScheduleId: scheduleID,
		SeatCount:  seatCount,
	})
}

func (c *BookingClient) GetBooking(ctx context.Context, bookingID int64) (*pb.GetBookingResponse, error) {
	return c.client.GetBooking(ctx, &pb.GetBookingRequest{
		BookingId: bookingID,
	})
}

func (c *BookingClient) ListUserBookings(ctx context.Context, userID int64, page, limit int32) (*pb.ListUserBookingsResponse, error) {
	return c.client.ListUserBookings(ctx, &pb.ListUserBookingsRequest{
		UserId: userID,
		Page:   page,
		Limit:  limit,
	})
}

func (c *BookingClient) CancelBooking(ctx context.Context, bookingID, userID int64) (*pb.CancelBookingResponse, error) {
	return c.client.CancelBooking(ctx, &pb.CancelBookingRequest{
		BookingId: bookingID,
		UserId:    userID,
	})
}

func (c *BookingClient) UpdatePaymentStatus(ctx context.Context, bookingId, userId int64, status string) (*pb.UpdatePaymentStatusResponse, error) {
	return c.client.UpdatePaymentStatus(ctx, &pb.UpdatePaymentStatusRequest{
		BookingId: bookingId,
		UserId:    userId,
		Status:    status,
	})
}
