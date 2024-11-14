package config_test

import (
	"agent/internal/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigWithFlags(t *testing.T) {
	os.Args = []string{
		"cmd",
		"-a=http://localhost:9090",
		"-p=10",
		"-r=30",
	}

	cfg, err := config.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:9090", cfg.ServerURL)
	assert.Equal(t, int64(10), cfg.PollInterval)
	assert.Equal(t, int64(30), cfg.ReportInterval)
}

func TestGetConfigWithEnvVariables(t *testing.T) {
	os.Setenv("SERVER_URL", "http://localhost:7070")
	os.Setenv("POLL_INTERVAL", "15")
	os.Setenv("REPORT_INTERVAL", "45")
	os.Setenv("SECRET_KEY", "supersecret")
	os.Setenv("LOG_LEVEL", "info")

	defer func() {
		os.Unsetenv("SERVER_URL")
		os.Unsetenv("POLL_INTERVAL")
		os.Unsetenv("REPORT_INTERVAL")
		os.Unsetenv("SECRET_KEY")
		os.Unsetenv("LOG_LEVEL")
	}()

	cfg, err := config.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:7070", cfg.ServerURL)
	assert.Equal(t, int64(15), cfg.PollInterval)
	assert.Equal(t, int64(45), cfg.ReportInterval)
	assert.Equal(t, "supersecret", cfg.SecretKey)
	assert.Equal(t, "info", cfg.LogLevel)
}

func TestGetConfigWithDefaults(t *testing.T) {
	os.Clearenv()

	cfg, err := config.GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080", cfg.ServerURL)
	assert.Equal(t, int64(5), cfg.PollInterval)
	assert.Equal(t, int64(20), cfg.ReportInterval)
	assert.Equal(t, "default", cfg.SecretKey)
	assert.Equal(t, "error", cfg.LogLevel)
}
