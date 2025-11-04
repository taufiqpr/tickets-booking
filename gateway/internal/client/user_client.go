package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "ticket-booking/proto/user"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(host string, port int) (*UserClient, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)
	return &UserClient{client: client}, nil
}

func (c *UserClient) Register(ctx context.Context, username, email, password, confirmPassword string) (*pb.RegisterResponse, error) {
	return c.client.Register(ctx, &pb.RegisterRequest{
		Username:        username,
		Email:           email,
		Password:        password,
		ConfirmPassword: confirmPassword,
	})
}

func (c *UserClient) Login(ctx context.Context, username, password string) (*pb.LoginResponse, error) {
	return c.client.Login(ctx, &pb.LoginRequest{
		Username: username,
		Password: password,
	})
}

func (c *UserClient) GetUser(ctx context.Context, userID int64) (*pb.GetUserResponse, error) {
	return c.client.GetUser(ctx, &pb.GetUserRequest{
		UserId: userID,
	})
}

func (c *UserClient) ForgotPassword(ctx context.Context, email string) (*pb.ForgotPasswordResponse, error) {
	return c.client.ForgotPassword(ctx, &pb.ForgotPasswordRequest{
		Email: email,
	})
}

func (c *UserClient) ResetPassword(ctx context.Context, token, newPassword string) (*pb.ResetPasswordResponse, error) {
	return c.client.ResetPassword(ctx, &pb.ResetPasswordRequest{
		Token:       token,
		NewPassword: newPassword,
	})
}
