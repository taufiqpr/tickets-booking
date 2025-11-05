package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "ticket-booking/proto/user"
	"ticket-booking/user-service/config"
	"ticket-booking/user-service/internal/handler"
	"ticket-booking/user-service/internal/repository"
	"ticket-booking/user-service/internal/routes"
	"ticket-booking/user-service/internal/service"
)

func main() {
	cfg, err := config.LoadEnv("USER_")
	if err != nil {
		log.Fatalf("load env: %v", err)
	}

	pool, err := newDBPool(cfg)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer pool.Close()

	log.Println("db connected")

	userRepo := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	httpHandler := handler.NewHTTPHandler(userService)
	grpcHandler := handler.NewGrpcServer(userService)

	httpMux := routes.SetupRoutes(httpHandler)

	go func() {
		log.Printf("HTTP server listening on %s", cfg.Addr())
		log.Fatal(http.ListenAndServe(cfg.Addr(), httpMux))
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, grpcHandler)

	reflection.Register(grpcServer)

	log.Printf("%s gRPC server listening on port %d", cfg.ServiceName, cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func newDBPool(cfg *config.Config) (*pgxpool.Pool, error) {
	url := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=" + cfg.DBSSLMode
	pcfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgxpool.NewWithConfig(ctx, pcfg)
	if err != nil {
		return nil, err
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	if err := pool.Ping(ctx2); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
