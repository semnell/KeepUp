package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/semnell/KeepUp/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterRoutes(t *testing.T) {
	router := gin.New()
	err := RegisterRoutes(router)
	require.NoError(t, err)

	// Perform a GET request to /metrics
	req, err := http.NewRequest("GET", "/metrics", nil)
	require.NoError(t, err)

	verifyMetricsEndpoint(t, router, req)
}

func verifyMetricsEndpoint(t *testing.T, router *gin.Engine, req *http.Request) {
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestInitPrometheusMetrics(t *testing.T) {
	// Create a temporary config file with a sample job
	tempFile := createTempConfigFile(t)
	defer os.Remove(tempFile.Name())

	// Load config from the temporary file
	conf := utils.LoadConfig(tempFile.Name())

	// Initialize Prometheus metrics
	err := InitPrometheusMetrics(conf)
	assert.NoError(t, err)
}

func TestPrometheusHandler(t *testing.T) {
	router := gin.New()
	router.GET("/metrics", prometheusHandler())

	// Perform a GET request to /metrics
	req, err := http.NewRequest("GET", "/metrics", nil)
	require.NoError(t, err)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpMarkerHandler(t *testing.T) {
	// Create a temporary config file with a sample job
	tempFile := createTempConfigFile(t)
	defer os.Remove(tempFile.Name())

	// Create a sample UpdateMetricPost instance
	updateMetric := utils.UpdateMetricPost{
		URL:          "http://example.com",
		Reason:       "Success",
		ResCode:      200,
		MarkUp:       true,
		ResponseTime: 10.5,
	}

	// Marshal the updateMetric to JSON
	updateMetricJSON, err := json.Marshal(updateMetric)
	require.NoError(t, err)

	// Perform a POST request to /callback
	req, err := http.NewRequest("POST", "/callback", bytes.NewBuffer(updateMetricJSON))
	require.NoError(t, err)

	resp := httptest.NewRecorder()
	router := gin.New()
	router.POST("/callback", upMarkerHandler())
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// Additional assertions can be added based on your specific requirements
}

// createTempConfigFile creates a temporary YAML config file with a sample job for testing
func createTempConfigFile(t *testing.T) *os.File {
	content := []byte(`
version: "1.0"
jobs:
  - name: testJob
    scheme: https
    url: example.com
    interval: 60
    timeout: 10
    headers:
      - key: Authorization
        value: Bearer Token
    method: GET
    expect:
      status: 200
      body: OK
`)

	tempFile, err := os.CreateTemp("", "test_config.yaml")
	require.NoError(t, err)
	defer tempFile.Close()

	_, err = tempFile.Write(content)
	require.NoError(t, err)

	return tempFile
}
