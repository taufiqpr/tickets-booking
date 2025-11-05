package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"ticket-booking/gateway/config"
	"ticket-booking/gateway/internal/middleware"
	"ticket-booking/gateway/internal/proxy"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("load env: %v", err)
	}

	// Initialize reverse proxy
	reverseProxy := proxy.NewReverseProxy(
		cfg.UserHost, cfg.UserPort,
		cfg.BookHost, cfg.BookPort,
		cfg.TrainHost, cfg.TrainPort,
		cfg.ScheduleHost, cfg.SchedulePort,
	)

	// Initialize middleware (for protected routes)
	authMiddleware := middleware.NewAuthMiddleware("your-jwt-secret-key-here")

	r := gin.Default()

	// Add logging middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Auth routes - direct proxy to user-service
	authGroup := r.Group("/api/auth")
	{
		authGroup.Any("/*path", reverseProxy.ProxyToUserService())
	}

	// Schedule routes - direct proxy to schedule-service
	scheduleGroup := r.Group("/api/schedules")
	{
		scheduleGroup.Any("/*path", reverseProxy.ProxyToScheduleService())
	}

	// Booking routes - with auth middleware + proxy to booking-service
	bookingGroup := r.Group("/api/bookings")
	bookingGroup.Use(authMiddleware.RequireAuth())
	{
		bookingGroup.Any("/*path", reverseProxy.ProxyToBookingService())
	}

	// Train routes - with auth middleware + proxy to train-service
	trainGroup := r.Group("/api/trains")
	trainGroup.Use(authMiddleware.RequireAuth())
	{
		trainGroup.Any("/*path", reverseProxy.ProxyToTrainService())
	}

	// Gateway health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "gateway",
			"message": "API Gateway is running",
		})
	})

	// Service health checks
	r.GET("/health/user", reverseProxy.HealthCheckProxy("user-service", reverseProxy.UserServiceURL))
	r.GET("/health/booking", reverseProxy.HealthCheckProxy("booking-service", reverseProxy.BookingServiceURL))
	r.GET("/health/train", reverseProxy.HealthCheckProxy("train-service", reverseProxy.TrainServiceURL))
	r.GET("/health/schedule", reverseProxy.HealthCheckProxy("schedule-service", reverseProxy.ScheduleServiceURL))

	log.Printf("🚀 Gateway (Reverse Proxy) listening on %s", cfg.Addr())
	log.Printf("📡 Proxying to services:")
	log.Printf("   - User Service: http://%s:%d", cfg.UserHost, cfg.UserPort)
	log.Printf("   - Booking Service: http://%s:%d", cfg.BookHost, cfg.BookPort)
	log.Printf("   - Train Service: http://%s:%d", cfg.TrainHost, cfg.TrainPort)
	log.Printf("   - Schedule Service: http://%s:%d", cfg.ScheduleHost, cfg.SchedulePort)

	log.Fatal(r.Run(cfg.Addr()))
}
