package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"ticket-booking/gateway/internal/client"
)

type ScheduleHandler struct {
	scheduleClient *client.ScheduleClient
}

func NewScheduleHandler(scheduleClient *client.ScheduleClient) *ScheduleHandler {
	return &ScheduleHandler{scheduleClient: scheduleClient}
}

func (h *ScheduleHandler) SearchSchedules(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("origin")
	destination := r.URL.Query().Get("destination")
	departureDate := r.URL.Query().Get("departure_date")
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

	resp, err := h.scheduleClient.ListSchedules(context.Background(), origin, destination, departureDate, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *ScheduleHandler) GetSchedule(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid schedule_id", http.StatusBadRequest)
		return
	}

	scheduleIdStr := parts[len(parts)-1]
	scheduleId, err := strconv.ParseInt(scheduleIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid schedule_id", http.StatusBadRequest)
		return
	}

	resp, err := h.scheduleClient.GetSchedule(context.Background(), scheduleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *ScheduleHandler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TrainId       int64   `json:"train_id"`
		Origin        string  `json:"origin"`
		Destination   string  `json:"destination"`
		DepartureTime string  `json:"departure_time"`
		ArrivalTime   string  `json:"arrival_time"`
		Price         float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.scheduleClient.CreateSchedule(context.Background(), req.TrainId, req.Origin, req.Destination, req.DepartureTime, req.ArrivalTime, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
