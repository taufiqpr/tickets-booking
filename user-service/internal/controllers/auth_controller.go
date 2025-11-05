package controllers

import (
	"encoding/json"
	"net/http"

	"ticket-booking/user-service/internal/helper"
	"ticket-booking/user-service/internal/models"
	"ticket-booking/user-service/internal/service"
)

type AuthController struct {
	userService service.UserService
}

func NewAuthController(userService service.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteBadRequest(w, "Invalid request body", err)
		return
	}

	user, token, err := c.userService.Register(r.Context(), req.Username, req.Email, req.Password, req.ConfirmPassword)
	if err != nil {
		helper.WriteBadRequest(w, "Registration failed", err)
		return
	}

	response := models.AuthResponse{
		User: &models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
		Token: token,
	}

	helper.WriteSuccess(w, "User registered successfully", response)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteBadRequest(w, "Invalid request body", err)
		return
	}

	user, token, err := c.userService.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		helper.WriteUnauthorized(w, "Login failed", err)
		return
	}

	response := models.AuthResponse{
		User: &models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
		Token: token,
	}

	helper.WriteSuccess(w, "Login successful", response)
}
