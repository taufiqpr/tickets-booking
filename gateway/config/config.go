package config

import "os"

type Config struct {
	Port         string
	UserBaseURL  string
	TrainBaseURL string
	SchedBaseURL string
	BookBaseURL  string
}

func Load() *Config {
	get := func(k, def string) string { if v := os.Getenv(k); v != "" { return v }; return def }
	return &Config{
		Port:         get("GATEWAY_PORT", "8080"),
		UserBaseURL:  get("USER_BASE_URL", "http://localhost:8081"),
		TrainBaseURL: get("TRAIN_BASE_URL", "http://localhost:8082"),
		SchedBaseURL: get("SCHEDULE_BASE_URL", "http://localhost:8083"),
		BookBaseURL:  get("BOOKING_BASE_URL", "http://localhost:8084"),
	}
}

func (c *Config) Addr() string { return ":" + c.Port }
