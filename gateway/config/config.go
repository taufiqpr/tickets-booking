package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	UserHost     string
	UserPort     int
	TrainHost    string
	TrainPort    int
	ScheduleHost string
	SchedulePort int
	BookHost     string
	BookPort     int
}

func LoadEnv() (*Config, error) {
	_ = godotenv.Load()

	getReqDefault := func(k, def string) string {
		if v := os.Getenv(k); v != "" {
			return v
		}
		return def
	}

	getPortDefault := func(k string, def int) int {
		if v := os.Getenv(k); v != "" {
			if port, err := strconv.Atoi(v); err == nil {
				return port
			}
		}
		return def
	}

	return &Config{
		Port:         getReqDefault("GATEWAY_PORT", "8080"),
		UserHost:     getReqDefault("USER_HOST", "localhost"),
		UserPort:     getPortDefault("USER_PORT", 8081),
		TrainHost:    getReqDefault("TRAIN_HOST", "localhost"),
		TrainPort:    getPortDefault("TRAIN_PORT", 8082),
		ScheduleHost: getReqDefault("SCHEDULE_HOST", "localhost"),
		SchedulePort: getPortDefault("SCHEDULE_PORT", 8083),
		BookHost:     getReqDefault("BOOKING_HOST", "localhost"),
		BookPort:     getPortDefault("BOOKING_PORT", 8084),
	}, nil
}

func (c *Config) Addr() string {
	return ":" + c.Port
}
