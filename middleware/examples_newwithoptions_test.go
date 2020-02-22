package middleware_test

import (
	"github.com/TsvetanMilanov/go-gin-prometheus-middleware/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func ExampleNewWithOptions_options() {
	options := &middleware.Options{
		// Overwrite the http metric name.
		HTTPMetricName: "my_request_duration_seconds",
		// Overwrite the http metric histogram buckets.
		HTTPMetricBuckets: []float64{0.1, 0.5, 1, 2, 5},
		// Don't use normalized paths for the path label.
		// IMPORTANT: If you have ids in the paths this can lead to huge amount of metrics.
		HTTPMetricUseNotNormalizedPaths: true,
		// Overwrite the metrics path. Metrics for the metrics path won't be collected.
		MetricsPath: "/my-metrics",
		// Add static labels to the http metrics histogram.
		AdditionalHTTPMetricDefaultLabelsNames: map[string]string{"custom": "my-value"},
		// Change the name of the default labels.
		HTTPMetricDefaultLabelsNames: map[string]string{middleware.StatusCodeLabelName: "status"},
	}

	middleware := middleware.NewWithOptions(options)

	gin.Default().Use(middleware)
}

func ExampleNew_registry() {
	myCustomRegistry := prometheus.NewRegistry()
	options := &middleware.Options{
		// Use custom registry for the metrics.
		Registry: myCustomRegistry,
	}

	middleware := middleware.NewWithOptions(options)

	gin.Default().Use(middleware)
}
