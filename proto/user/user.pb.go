package user

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterRequest represents registration request
type RegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (*RegisterRequest) ProtoMessage()    {}

func (m *RegisterRequest) GetUsername() string        { return m.Username }
func (m *RegisterRequest) GetEmail() string           { return m.Email }
func (m *RegisterRequest) GetPassword() string        { return m.Password }
func (m *RegisterRequest) GetConfirmPassword() string { return m.ConfirmPassword }

// RegisterResponse represents registration response
type RegisterResponse struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (*RegisterResponse) ProtoMessage()    {}

func (m *RegisterResponse) GetUserId() int64    { return m.UserId }
func (m *RegisterResponse) GetUsername() string { return m.Username }
func (m *RegisterResponse) GetEmail() string    { return m.Email }
func (m *RegisterResponse) GetToken() string    { return m.Token }

// LoginRequest represents login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (*LoginRequest) ProtoMessage()    {}

func (m *LoginRequest) GetUsername() string { return m.Username }
func (m *LoginRequest) GetPassword() string { return m.Password }

// LoginResponse represents login response
type LoginResponse struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (*LoginResponse) ProtoMessage()    {}

func (m *LoginResponse) GetUserId() int64    { return m.UserId }
func (m *LoginResponse) GetUsername() string { return m.Username }
func (m *LoginResponse) GetEmail() string    { return m.Email }
func (m *LoginResponse) GetToken() string    { return m.Token }

// GetUserRequest represents get user request
type GetUserRequest struct {
	UserId int64 `json:"user_id"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (*GetUserRequest) ProtoMessage()    {}

func (m *GetUserRequest) GetUserId() int64 { return m.UserId }

// GetUserResponse represents get user response
type GetUserResponse struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (*GetUserResponse) ProtoMessage()    {}

func (m *GetUserResponse) GetUserId() int64    { return m.UserId }
func (m *GetUserResponse) GetUsername() string { return m.Username }
func (m *GetUserResponse) GetEmail() string    { return m.Email }

// ForgotPasswordRequest represents forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

func (m *ForgotPasswordRequest) Reset()         { *m = ForgotPasswordRequest{} }
func (m *ForgotPasswordRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (*ForgotPasswordRequest) ProtoMessage()    {}

func (m *ForgotPasswordRequest) GetEmail() string { return m.Email }

// ForgotPasswordResponse represents forgot password response
type ForgotPasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (m *ForgotPasswordResponse) Reset()         { *m = ForgotPasswordResponse{} }
func (m *ForgotPasswordResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (*ForgotPasswordResponse) ProtoMessage()    {}

func (m *ForgotPasswordResponse) GetSuccess() bool   { return m.Success }
func (m *ForgotPasswordResponse) GetMessage() string { return m.Message }
func (m *ForgotPasswordResponse) GetToken() string   { return m.Token }

// ResetPasswordRequest represents reset password request
type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func (m *ResetPasswordRequest) Reset()         { *m = ResetPasswordRequest{} }
func (m *ResetPasswordRequest) String() string { return fmt.Sprintf("%+v", *m) }
func (*ResetPasswordRequest) ProtoMessage()    {}

func (m *ResetPasswordRequest) GetToken() string       { return m.Token }
func (m *ResetPasswordRequest) GetNewPassword() string { return m.NewPassword }

// ResetPasswordResponse represents reset password response
type ResetPasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (m *ResetPasswordResponse) Reset()         { *m = ResetPasswordResponse{} }
func (m *ResetPasswordResponse) String() string { return fmt.Sprintf("%+v", *m) }
func (*ResetPasswordResponse) ProtoMessage()    {}

func (m *ResetPasswordResponse) GetSuccess() bool   { return m.Success }
func (m *ResetPasswordResponse) GetMessage() string { return m.Message }

// UserServiceClient is the client API for UserService service.
type UserServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}
func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _UserService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _UserService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
