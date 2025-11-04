package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"ticket-booking/gateway/internal/client"
)

type BookingHandler struct {
	bookingClient *client.BookingClient
}

func NewBookingHandler(bookingClient *client.BookingClient) *BookingHandler {
	return &BookingHandler{bookingClient: bookingClient}
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserId     int64 `json:"user_id"`
		ScheduleId int64 `json:"schedule_id"`
		SeatCount  int32 `json:"seat_count"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.bookingClient.CreateBooking(context.Background(), req.UserId, req.ScheduleId, req.SeatCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	// Extract booking ID from URL path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid booking_id", http.StatusBadRequest)
		return
	}

	bookingIdStr := parts[len(parts)-1]
	bookingId, err := strconv.ParseInt(bookingIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid booking_id", http.StatusBadRequest)
		return
	}

	resp, err := h.bookingClient.GetBooking(context.Background(), bookingId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *BookingHandler) ListUserBookings(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	userIdStr := parts[len(parts)-1]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := int32(1)
	limit := int32(10)

	if pageStr != "" {
		if p, err := strconv.ParseInt(pageStr, 10, 32); err == nil {
			page = int32(p)
		}
	}

	if limitStr != "" {
		if l, err := strconv.ParseInt(limitStr, 10, 32); err == nil {
			limit = int32(l)
		}
	}

	resp, err := h.bookingClient.ListUserBookings(context.Background(), userId, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *BookingHandler) UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BookingId int64  `json:"booking_id"`
		UserId    int64  `json:"user_id"`
		Status    string `json:"status"` // success, failed, expired
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.bookingClient.UpdatePaymentStatus(context.Background(), req.BookingId, req.UserId, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	// Extract booking ID from URL path
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid booking_id", http.StatusBadRequest)
		return
	}

	bookingIdStr := parts[len(parts)-1]
	bookingId, err := strconv.ParseInt(bookingIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid booking_id", http.StatusBadRequest)
		return
	}

	var req struct {
		UserId int64 `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.bookingClient.CancelBooking(context.Background(), bookingId, req.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
