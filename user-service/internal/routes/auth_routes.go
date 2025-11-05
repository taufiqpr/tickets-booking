package routes

import (
	"net/http"

	"ticket-booking/user-service/internal/handler"
)

func SetupAuthRoutes(mux *http.ServeMux, httpHandler *handler.HTTPHandler) {
	mux.HandleFunc("/auth/register", httpHandler.Register)
	mux.HandleFunc("/auth/login", httpHandler.Login)
	mux.HandleFunc("/auth/forgot-password", httpHandler.ForgotPassword)
	mux.HandleFunc("/auth/reset-password", httpHandler.ResetPassword)
}
