// Simple proto implementation for booking service
package booking

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Booking represents a booking
type Booking struct {
	Id            int64   `json:"id"`
	UserId        int64   `json:"user_id"`
	ScheduleId    int64   `json:"schedule_id"`
	BookingCode   string  `json:"booking_code"`
	Status        string  `json:"status"`
	TotalPrice    float64 `json:"total_price"`
	SeatCount     int32   `json:"seat_count"`
	CreatedAt     string  `json:"created_at"`
	ExpiresAt     string  `json:"expires_at"`
	Origin        string  `json:"origin"`
	Destination   string  `json:"destination"`
	DepartureTime string  `json:"departure_time"`
	ArrivalTime   string  `json:"arrival_time"`
	TrainName     string  `json:"train_name"`
}

// CreateBookingRequest represents create booking request
type CreateBookingRequest struct {
	UserId     int64 `json:"user_id"`
	ScheduleId int64 `json:"schedule_id"`
	SeatCount  int32 `json:"seat_count"`
}

// CreateBookingResponse represents create booking response
type CreateBookingResponse struct {
	Booking *Booking `json:"booking"`
}

// GetBookingRequest represents get booking request
type GetBookingRequest struct {
	BookingId int64 `json:"booking_id"`
}

// GetBookingResponse represents get booking response
type GetBookingResponse struct {
	Booking *Booking `json:"booking"`
}

// ListUserBookingsRequest represents list user bookings request
type ListUserBookingsRequest struct {
	UserId int64 `json:"user_id"`
	Page   int32 `json:"page"`
	Limit  int32 `json:"limit"`
}

// ListUserBookingsResponse represents list user bookings response
type ListUserBookingsResponse struct {
	Bookings []*Booking `json:"bookings"`
	Total    int32      `json:"total"`
}

// CancelBookingRequest represents cancel booking request
type CancelBookingRequest struct {
	BookingId int64 `json:"booking_id"`
	UserId    int64 `json:"user_id"`
}

// CancelBookingResponse represents cancel booking response
type CancelBookingResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UpdatePaymentStatusRequest represents update payment status request
type UpdatePaymentStatusRequest struct {
	BookingId int64  `json:"booking_id"`
	UserId    int64  `json:"user_id"`
	Status    string `json:"status"`
}

// UpdatePaymentStatusResponse represents update payment status response
type UpdatePaymentStatusResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Booking *Booking `json:"booking"`
}

// BookingServiceClient is the client API for BookingService service.
type BookingServiceClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GetBooking(ctx context.Context, in *GetBookingRequest, opts ...grpc.CallOption) (*GetBookingResponse, error)
	ListUserBookings(ctx context.Context, in *ListUserBookingsRequest, opts ...grpc.CallOption) (*ListUserBookingsResponse, error)
	CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error)
	UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBooking(ctx context.Context, in *GetBookingRequest, opts ...grpc.CallOption) (*GetBookingResponse, error) {
	out := new(GetBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) ListUserBookings(ctx context.Context, in *ListUserBookingsRequest, opts ...grpc.CallOption) (*ListUserBookingsResponse, error) {
	out := new(ListUserBookingsResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/ListUserBookings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CancelBooking(ctx context.Context, in *CancelBookingRequest, opts ...grpc.CallOption) (*CancelBookingResponse, error) {
	out := new(CancelBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error) {
	out := new(UpdatePaymentStatusResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/UpdatePaymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
type BookingServiceServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GetBooking(context.Context, *GetBookingRequest) (*GetBookingResponse, error)
	ListUserBookings(context.Context, *ListUserBookingsRequest) (*ListUserBookingsResponse, error)
	CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error)
	UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error)
}

// UnimplementedBookingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (*UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}

func (*UnimplementedBookingServiceServer) GetBooking(context.Context, *GetBookingRequest) (*GetBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooking not implemented")
}

func (*UnimplementedBookingServiceServer) ListUserBookings(context.Context, *ListUserBookingsRequest) (*ListUserBookingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserBookings not implemented")
}

func (*UnimplementedBookingServiceServer) CancelBooking(context.Context, *CancelBookingRequest) (*CancelBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}

func (*UnimplementedBookingServiceServer) UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentStatus not implemented")
}

func RegisterBookingServiceServer(s *grpc.Server, srv BookingServiceServer) {
	s.RegisterService(&_BookingService_serviceDesc, srv)
}

var _BookingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GetBooking",
			Handler:    _BookingService_GetBooking_Handler,
		},
		{
			MethodName: "ListUserBookings",
			Handler:    _BookingService_ListUserBookings_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _BookingService_CancelBooking_Handler,
		},
		{
			MethodName: "UpdatePaymentStatus",
			Handler:    _BookingService_UpdatePaymentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBooking(ctx, req.(*GetBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_ListUserBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserBookingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ListUserBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/ListUserBookings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ListUserBookings(ctx, req.(*ListUserBookingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CancelBooking(ctx, req.(*CancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_UpdatePaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).UpdatePaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/UpdatePaymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).UpdatePaymentStatus(ctx, req.(*UpdatePaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}
