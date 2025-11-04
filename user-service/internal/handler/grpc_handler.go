package handler

import (
	"context"

	pb "ticket-booking/proto/user"
	"ticket-booking/user-service/internal/service"
)

type GrpcServer struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewGrpcServer(userService service.UserService) *GrpcServer {
	return &GrpcServer{
		userService: userService,
	}
}

func (s *GrpcServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, token, err := s.userService.Register(ctx, req.Username, req.Email, req.Password, req.FullName)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		Token:    token,
	}, nil
}

func (s *GrpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, token, err := s.userService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		Token:    token,
	}, nil
}

func (s *GrpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userService.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	}, nil
}

func (s *GrpcServer) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	token, err := s.userService.ForgotPassword(ctx, req.Email)
	if err != nil {
		return &pb.ForgotPasswordResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.ForgotPasswordResponse{
		Success: true,
		Token:   token,
		Message: "Reset token sent to email",
	}, nil
}

func (s *GrpcServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	err := s.userService.ResetPassword(ctx, req.Token, req.NewPassword)
	if err != nil {
		return &pb.ResetPasswordResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.ResetPasswordResponse{
		Success: true,
		Message: "Password reset successful",
	}, nil
}
