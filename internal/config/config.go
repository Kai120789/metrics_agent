package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	ServerAddress  string
	LogLevel       string
	PollInterval   string
	ReportInterval string
}

func GetConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}

	flag.StringVar(&cfg.ServerAddress, "a", "localhost:8080", "address and port to run server")

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		cfg.ServerAddress = envRunAddr
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		cfg.LogLevel = envLogLevel
	} else {
		cfg.LogLevel = zapcore.ErrorLevel.String()
	}

	flag.StringVar(&cfg.PollInterval, "p", "5", "poll interval (sec)")

	if envPollInt := os.Getenv("POLL_INTERVAL"); envPollInt != "" {
		cfg.PollInterval = envPollInt
	}

	flag.StringVar(&cfg.ReportInterval, "r", "20", "report interval (sec)")

	if envRepInt := os.Getenv("REPORT_INTERVAL"); envRepInt != "" {
		cfg.ReportInterval = envRepInt
	}

	flag.Parse()

	return cfg, nil
}
