package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string
	Port        string
	GRPCPort    int
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPassword  string
	DBSSLMode   string
	JWTSecret   string
}

func LoadEnv(prefix string) (*Config, error) {
	_ = godotenv.Load()

	getReq := func(k string) (string, error) {
		if v := os.Getenv(prefix + k); v != "" {
			return v, nil
		}
		if v := os.Getenv(k); v != "" {
			return v, nil
		}
		return "", fmt.Errorf("missing env %s or %s", prefix+k, k)
	}

	name, err := getReq("SERVICE_NAME")
	if err != nil {
		return nil, err
	}
	port, err := getReq("PORT")
	if err != nil {
		return nil, err
	}
	grpcPortStr, err := getReq("GRPC_PORT")
	if err != nil {
		return nil, err
	}
	grpcPort, err := strconv.Atoi(grpcPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid GRPC_PORT: %v", err)
	}
	host, err := getReq("DB_HOST")
	if err != nil {
		return nil, err
	}
	dbPort, err := getReq("DB_PORT")
	if err != nil {
		return nil, err
	}
	dbName, err := getReq("DB_NAME")
	if err != nil {
		return nil, err
	}
	dbUser, err := getReq("DB_USER")
	if err != nil {
		return nil, err
	}
	dbPass, err := getReq("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	sslMode, err := getReq("DB_SSL_MODE")
	if err != nil {
		return nil, err
	}
	jwtSecret, err := getReq("JWT_SECRET")
	if err != nil {
		return nil, err
	}

	return &Config{
		ServiceName: name,
		Port:        port,
		GRPCPort:    grpcPort,
		DBHost:      host,
		DBPort:      dbPort,
		DBName:      dbName,
		DBUser:      dbUser,
		DBPassword:  dbPass,
		DBSSLMode:   sslMode,
		JWTSecret:   jwtSecret,
	}, nil
}

func (c *Config) Addr() string { return ":" + c.Port }
