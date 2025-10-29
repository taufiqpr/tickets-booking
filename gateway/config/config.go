package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	UserBaseURL  string
	TrainBaseURL string
	SchedBaseURL string
	BookBaseURL  string
}

func LoadEnv() (*Config, error) {
	_ = godotenv.Load()

	getReq := func(k string) (string, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}
		return "", fmt.Errorf("missing env %s", k)
	}

	port, err := getReq("GATEWAY_PORT")
	if err != nil { return nil, err }
	user, err := getReq("USER_BASE_URL")
	if err != nil { return nil, err }
	train, err := getReq("TRAIN_BASE_URL")
	if err != nil { return nil, err }
	sched, err := getReq("SCHEDULE_BASE_URL")
	if err != nil { return nil, err }
	book, err := getReq("BOOKING_BASE_URL")
	if err != nil { return nil, err }

	return &Config{
		Port:         port,
		UserBaseURL:  user,
		TrainBaseURL: train,
		SchedBaseURL: sched,
		BookBaseURL:  book,
	}, nil
}

func (c *Config) Addr() string { return ":" + c.Port }
