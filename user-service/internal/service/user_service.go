package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"ticket-booking/user-service/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, username, email, password, fullName string) (*repository.User, string, error)
	Login(ctx context.Context, username, password string) (*repository.User, string, error)
	GetUser(ctx context.Context, id int64) (*repository.User, error)
	ForgotPassword(ctx context.Context, email string) (string, error)
	ResetPassword(ctx context.Context, token, newPassword string) error
}

type userService struct {
	userRepo repository.UserRepository
	jwtKey   []byte
}

func NewUserService(userRepo repository.UserRepository, jwtKey string) UserService {
	return &userService{
		userRepo: userRepo,
		jwtKey:   []byte(jwtKey),
	}
}

func (s *userService) Register(ctx context.Context, username, email, password, fullName string) (*repository.User, string, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetByUsername(ctx, username)
	if existingUser != nil {
		return nil, "", errors.New("username already exists")
	}

	existingUser, _ = s.userRepo.GetByEmail(ctx, email)
	if existingUser != nil {
		return nil, "", errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &repository.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		FullName: fullName,
	}

	user, err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *userService) Login(ctx context.Context, username, password string) (*repository.User, string, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*repository.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) ForgotPassword(ctx context.Context, email string) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("email not found")
	}

	// Generate reset token (simplified - in production use proper token generation)
	resetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
		"type":    "reset",
	})

	tokenString, err := resetToken.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	// In production, send email with reset link
	// For now, just return the token
	return tokenString, nil
}

func (s *userService) ResetPassword(ctx context.Context, token, newPassword string) error {
	// Parse and validate token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return errors.New("invalid or expired token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "reset" {
		return errors.New("invalid token type")
	}

	userID := int64(claims["user_id"].(float64))

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password
	return s.userRepo.UpdatePassword(ctx, userID, string(hashedPassword))
}

func (s *userService) generateToken(user *repository.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
