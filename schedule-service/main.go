package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"ticket-booking/schedule-service/config"
)

func main() {
	cfg, err := config.LoadEnv("SCHEDULE_")
	if err != nil {
		log.Fatalf("load env: %v", err)
	}

	pool, err := newDBPool(cfg)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer pool.Close()

	_ = runInitUp(pool)
	log.Println("db connected")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); _, _ = w.Write([]byte("ok")) })
	log.Printf("%s listening on %s", cfg.ServiceName, cfg.Addr())
	log.Fatal(http.ListenAndServe(cfg.Addr(), mux))
}

func newDBPool(cfg *config.Config) (*pgxpool.Pool, error) {
	url := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=" + cfg.DBSSLMode
	pcfg, err := pgxpool.ParseConfig(url)
	if err != nil { return nil, err }
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgxpool.NewWithConfig(ctx, pcfg)
	if err != nil { return nil, err }
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	if err := pool.Ping(ctx2); err != nil { pool.Close(); return nil, err }
	return pool, nil
}

func runInitUp(pool *pgxpool.Pool) error {
	b, err := os.ReadFile("db/migration/init_up.sql")
	if err != nil { return nil }
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = pool.Exec(ctx, string(b))
	return err
}
