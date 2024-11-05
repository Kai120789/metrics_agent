package api_test

import (
	"agent/internal/api"
	"agent/internal/dto"
	"agent/internal/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendMetrics(t *testing.T) {
	var metrics [31]dto.Metric
	metrics[0] = dto.Metric{Name: "PollCount", Type: "counter", Delta: new(int64)}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "/api/updates", r.URL.Path)

		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, utils.GenerateHash("test_key"), r.Header.Get("Hash"))

		var receivedMetrics [31]dto.Metric
		err := json.NewDecoder(r.Body).Decode(&receivedMetrics)
		require.NoError(t, err)
		assert.Equal(t, metrics, receivedMetrics)

		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	apiClient := api.New()
	err := apiClient.SendMetrics(metrics, server.URL, "test_key")

	require.NoError(t, err)
}
