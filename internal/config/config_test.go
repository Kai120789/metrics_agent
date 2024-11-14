package config_test

import (
	"agent/internal/config"
	"os"
	"strconv"
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

func TestGetEnvStringOrDefault(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, "default_value", getEnvStringOrDefault("NON_EXISTENT", "default_value"))

	os.Setenv("EXISTENT_VAR", "set_value")
	defer os.Unsetenv("EXISTENT_VAR")
	assert.Equal(t, "set_value", getEnvStringOrDefault("EXISTENT_VAR", "default_value"))
}

func TestGetEnvIntOrDefault(t *testing.T) {
	os.Clearenv()

	intVal, err := getEnvIntOrDefault("NON_EXISTENT", 10)
	assert.NoError(t, err)
	assert.Equal(t, 10, *intVal)

	os.Setenv("EXISTENT_INT", "25")
	defer os.Unsetenv("EXISTENT_INT")
	intVal, err = getEnvIntOrDefault("EXISTENT_INT", 10)
	assert.NoError(t, err)
	assert.Equal(t, 25, *intVal)

	os.Setenv("INVALID_INT", "invalid")
	defer os.Unsetenv("INVALID_INT")
	intVal, err = getEnvIntOrDefault("INVALID_INT", 10)
	assert.Error(t, err)
	assert.Nil(t, intVal)
}

func getEnvStringOrDefault(name, defaultValue string) string {
	if envString := os.Getenv(name); envString != "" {
		return envString
	}

	return defaultValue
}

func getEnvIntOrDefault(name string, defaultValue int) (*int, error) {
	if envInt := os.Getenv(name); envInt != "" {
		intEnvInt, err := strconv.Atoi(envInt)
		if err != nil {
			return nil, err
		}
		return &intEnvInt, nil
	}

	return &defaultValue, nil
}
