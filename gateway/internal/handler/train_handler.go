package handler

import (
	"context"
	"net/http"
	"strconv"

	"ticket-booking/gateway/internal/client"

	"github.com/gin-gonic/gin"
)

type TrainHandler struct {
	trainClient *client.TrainClient
}

func NewTrainHandler(trainClient *client.TrainClient) *TrainHandler {
	return &TrainHandler{trainClient: trainClient}
}

func (h *TrainHandler) CreateTrain(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Capacity int32  `json:"capacity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.trainClient.CreateTrain(context.Background(), req.Name, req.Type, req.Capacity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *TrainHandler) GetTrain(c *gin.Context) {
	trainIdStr := c.Param("id")
	trainId, err := strconv.ParseInt(trainIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid train_id"})
		return
	}

	resp, err := h.trainClient.GetTrain(context.Background(), trainId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *TrainHandler) ListTrains(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		page = 1
	}

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		limit = 10
	}

	resp, err := h.trainClient.ListTrains(context.Background(), int32(page), int32(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *TrainHandler) UpdateTrain(c *gin.Context) {
	trainIdStr := c.Param("id")
	trainId, err := strconv.ParseInt(trainIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid train_id"})
		return
	}

	var req struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Capacity int32  `json:"capacity"`
		Status   string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.trainClient.UpdateTrain(context.Background(), trainId, req.Name, req.Type, req.Status, req.Capacity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *TrainHandler) DeleteTrain(c *gin.Context) {
	trainIdStr := c.Param("id")
	trainId, err := strconv.ParseInt(trainIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid train_id"})
		return
	}

	resp, err := h.trainClient.DeleteTrain(context.Background(), trainId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
