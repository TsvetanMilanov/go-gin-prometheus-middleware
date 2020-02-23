package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	prometheus_go "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

const (
	defaultLabelsLength = 3
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func setOKResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "OK",
	})
}

func assertLabel(t *testing.T, labels []*prometheus_go.LabelPair, label, expected string) {
	for _, pair := range labels {
		if pair.GetName() == label {
			assert.Equal(t, pair.GetValue(), expected)
			return
		}
	}

	assert.Fail(t, fmt.Sprintf("label '%s' not found", label))
}

func assertHistogramMetric(t *testing.T, metric *prometheus_go.Metric, expectedLabelsLength int, expectedMethod, expectedPath, expectedStatus string) {
	assert.NotNil(t, metric.Histogram)
	assert.Len(t, metric.Label, expectedLabelsLength, "Metric labels length mismatch")
	assertLabel(t, metric.Label, "method", expectedMethod)
	assertLabel(t, metric.Label, "path", expectedPath)
	assertLabel(t, metric.Label, "status_code", expectedStatus)
}

func createRouterAndRegistry(options ...*Options) (*gin.Engine, *prometheus.Registry) {
	router := gin.New()

	registry := prometheus.NewRegistry()
	opts := new(Options)
	if len(options) > 0 {
		opts = options[0]
	}

	opts.Registry = registry

	router.Use(NewWithOptions(opts))

	router.GET("/param/value/:param", setOKResponse)
	router.GET("/metrics", setOKResponse)
	router.GET("/", setOKResponse)

	return router, registry
}

func TestMiddleware(t *testing.T) {
	router, registry := createRouterAndRegistry()

	performRequest(router, http.MethodGet, "/")
	performRequest(router, http.MethodGet, "/param/value/value")

	records, err := registry.Gather()

	assert.NoError(t, err, "Gather should not throw error")
	assert.Len(t, records, 1, "Gathered metrics length mismatch")

	record := records[0]
	assert.Equal(t, "http_request_duration_seconds", *record.Name)

	assert.Len(t, record.Metric, 2, "Metrics length mismatch")

	assertHistogramMetric(t, record.Metric[0], defaultLabelsLength, http.MethodGet, "/", "200")
	assertHistogramMetric(t, record.Metric[1], defaultLabelsLength, http.MethodGet, "/param/value/:param", "200")
}

func TestMiddlewareNotNormalizedPaths(t *testing.T) {
	opts := &Options{HTTPMetricUseNotNormalizedPaths: true}
	router, registry := createRouterAndRegistry(opts)

	performRequest(router, http.MethodGet, "/param/value/value")

	records, err := registry.Gather()

	assert.NoError(t, err, "Gather should not throw error")
	assert.Len(t, records, 1, "Gathered metrics length mismatch")

	record := records[0]
	assert.Equal(t, "http_request_duration_seconds", *record.Name)

	assert.Len(t, record.Metric, 1, "Metrics length mismatch")

	assertHistogramMetric(t, record.Metric[0], defaultLabelsLength, http.MethodGet, "/param/value/value", "200")
}

func TestMiddlewareMetricsEndpoint(t *testing.T) {
	opts := &Options{HTTPMetricUseNotNormalizedPaths: true}
	router, registry := createRouterAndRegistry(opts)

	performRequest(router, http.MethodGet, "/metrics")

	records, err := registry.Gather()

	assert.NoError(t, err, "Gather should not throw error")
	assert.Len(t, records, 0, "Gathered metrics length mismatch")
}

func TestMiddlewareCustomLabels(t *testing.T) {
	opts := &Options{
		AdditionalHTTPMetricDefaultLabelsNames: map[string]string{"custom": "test_value"},
	}
	router, registry := createRouterAndRegistry(opts)

	performRequest(router, http.MethodGet, "/")

	records, err := registry.Gather()

	assert.NoError(t, err, "Gather should not throw error")
	assert.Len(t, records, 1, "Gathered metrics length mismatch")

	record := records[0]
	assert.Equal(t, "http_request_duration_seconds", *record.Name)

	assert.Len(t, record.Metric, 1, "Metrics length mismatch")

	metric := record.Metric[0]
	assertHistogramMetric(t, metric, defaultLabelsLength+1, http.MethodGet, "/", "200")
	assertLabel(t, metric.Label, "custom", "test_value")
}

func TestMiddlewareBlacklister(t *testing.T) {
	blacklister := func(c *gin.Context) bool {
		if c.Request.URL.Path == "/blacklist" {
			return true
		}

		return false
	}

	opts := &Options{Blacklister: blacklister}
	router, registry := createRouterAndRegistry(opts)

	performRequest(router, http.MethodGet, "/blacklist")
	performRequest(router, http.MethodGet, "/")

	records, err := registry.Gather()

	assert.NoError(t, err, "Gather should not throw error")
	assert.Len(t, records, 1, "Gathered metrics length mismatch")

	record := records[0]
	assert.Equal(t, "http_request_duration_seconds", *record.Name)

	assert.Len(t, record.Metric, 1, "Metrics length mismatch")

	assertHistogramMetric(t, record.Metric[0], defaultLabelsLength, http.MethodGet, "/", "200")
}
