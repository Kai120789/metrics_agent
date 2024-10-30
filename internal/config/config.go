package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	ServerURL      string
	LogLevel       string
	PollInterval   int64
	ReportInterval int64
	SecretKey      string
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}

	flag.StringVar(&cfg.ServerURL, "a", "http://localhost:8080", "URL and port to run server")
	flag.Int64Var(&cfg.PollInterval, "p", 5, "poll interval (sec)")
	flag.Int64Var(&cfg.ReportInterval, "r", 20, "report interval (sec)")

	cfg.SecretKey = getEnvStringOrDefault("SECRET_KEY", "default")
	cfg.ServerURL = getEnvStringOrDefault("SERVER_URL", "http://localhost:8080")
	pollInt, err := getEnvIntOrDefault("POLL_INTERVAL", 5)
	if err != nil {
		return nil, err
	}

	repInt, err := getEnvIntOrDefault("REPORT_INTERVAL", 20)
	if err != nil {
		return nil, err
	}

	cfg.PollInterval = int64(*pollInt)
	cfg.ReportInterval = int64(*repInt)

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		cfg.LogLevel = envLogLevel
	} else {
		cfg.LogLevel = zapcore.ErrorLevel.String()
	}

	flag.Parse()

	return cfg, nil
}

func getEnvStringOrDefault(name, defaultValue string) string {
	if envRunAddr := os.Getenv(name); envRunAddr != "" {
		return envRunAddr
	}

	return defaultValue
}

func getEnvIntOrDefault(name string, defaultValue int) (*int, error) {
	if envRepInt := os.Getenv(name); envRepInt != "" {
		intEnvRepInt, err := strconv.Atoi(envRepInt)
		if err != nil {
			return nil, err
		}
		return &intEnvRepInt, nil
	}

	return &defaultValue, nil
}
