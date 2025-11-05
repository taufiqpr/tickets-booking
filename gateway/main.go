package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"ticket-booking/gateway/config"
	"ticket-booking/gateway/internal/client"
	"ticket-booking/gateway/internal/handler"
	"ticket-booking/gateway/internal/middleware"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("load env: %v", err)
	}

	userClient := client.NewUserHTTPClient(cfg.UserHost, cfg.UserPort)

	bookingClient, err := client.NewBookingClient(cfg.BookHost, cfg.BookPort)
	if err != nil {
		log.Fatalf("create booking client: %v", err)
	}

	trainClient, err := client.NewTrainClient(cfg.TrainHost, cfg.TrainPort)
	if err != nil {
		log.Fatalf("create train client: %v", err)
	}

	scheduleClient, err := client.NewScheduleClient(cfg.ScheduleHost, cfg.SchedulePort)
	if err != nil {
		log.Fatalf("create schedule client: %v", err)
	}

	authMiddleware := middleware.NewAuthMiddleware("your-jwt-secret-key-here")

	authHandler := handler.NewAuthHandler(userClient)
	bookingHandler := handler.NewBookingHandler(bookingClient)
	trainHandler := handler.NewTrainHandler(trainClient)
	scheduleHandler := handler.NewScheduleHandler(scheduleClient)

	r := gin.Default()

	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", gin.WrapF(authHandler.Register))
		authGroup.POST("/login", gin.WrapF(authHandler.Login))
		authGroup.POST("/forgot-password", gin.WrapF(authHandler.ForgotPassword))
		authGroup.POST("/reset-password", gin.WrapF(authHandler.ResetPassword))
	}

	scheduleGroup := r.Group("/api/schedules")
	{
		scheduleGroup.GET("/search", gin.WrapF(scheduleHandler.SearchSchedules))
		scheduleGroup.GET("/:id", gin.WrapF(scheduleHandler.GetSchedule))
		scheduleGroup.POST("/", gin.WrapF(scheduleHandler.CreateSchedule))
	}

	bookingGroup := r.Group("/api/bookings")
	bookingGroup.Use(authMiddleware.RequireAuth())
	{
		bookingGroup.POST("/create", gin.WrapF(bookingHandler.CreateBooking))
		bookingGroup.GET("/:id", gin.WrapF(bookingHandler.GetBooking))
		bookingGroup.GET("/user/:userId", gin.WrapF(bookingHandler.ListUserBookings))
		bookingGroup.PUT("/payment-status", gin.WrapF(bookingHandler.UpdatePaymentStatus))
		bookingGroup.DELETE("/:id", gin.WrapF(bookingHandler.CancelBooking))
	}

	trainGroup := r.Group("/api/trains")
	trainGroup.Use(authMiddleware.RequireAuth())
	{
		trainGroup.POST("/", trainHandler.CreateTrain)
		trainGroup.GET("/:id", trainHandler.GetTrain)
		trainGroup.GET("/", trainHandler.ListTrains)
		trainGroup.PUT("/:id", trainHandler.UpdateTrain)
		trainGroup.DELETE("/:id", trainHandler.DeleteTrain)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("gateway listening on %s", cfg.Addr())
	log.Fatal(r.Run(cfg.Addr()))
}
