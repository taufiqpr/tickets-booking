package config

import "os"

type Config struct {
	ServiceName string
	Port        string
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPassword  string
	DBSSLMode   string
}

func LoadConfigFromEnv(prefix, defaultName, defaultPort string) *Config {
	get := func(k, def string) string {
		if v := os.Getenv(prefix + k); v != "" { return v }
		if v := os.Getenv(k); v != "" { return v }
		return def
	}
	return &Config{
		ServiceName: get("SERVICE_NAME", defaultName),
		Port:        get("PORT", defaultPort),
		DBHost:      get("DB_HOST", "localhost"),
		DBPort:      get("DB_PORT", "5432"),
		DBName:      get("DB_NAME", ""),
		DBUser:      get("DB_USER", ""),
		DBPassword:  get("DB_PASSWORD", ""),
		DBSSLMode:   get("DB_SSLMODE", "disable"),
	}
}

func (c *Config) Addr() string { return ":" + c.Port }
