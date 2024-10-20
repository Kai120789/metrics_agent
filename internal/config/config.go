package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	ServerAddress  string
	LogLevel       string
	PollInterval   int64
	ReportInterval int64
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

	flag.Int64Var(&cfg.PollInterval, "p", 5, "poll interval (sec)")

	if envPollInt := os.Getenv("POLL_INTERVAL"); envPollInt != "" {
		intEnvPollInt, err := strconv.Atoi(envPollInt)
		if err != nil {
			return nil, err
		}
		cfg.PollInterval = int64(intEnvPollInt)
	}

	flag.Int64Var(&cfg.ReportInterval, "r", 20, "report interval (sec)")

	if envRepInt := os.Getenv("REPORT_INTERVAL"); envRepInt != "" {
		intEnvRepInt, err := strconv.Atoi(envRepInt)
		if err != nil {
			return nil, err
		}
		cfg.ReportInterval = int64(intEnvRepInt)
	}

	flag.Parse()

	return cfg, nil
}
