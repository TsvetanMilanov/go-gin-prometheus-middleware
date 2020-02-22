package middleware

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionsSetHTTPMetricName(t *testing.T) {
	opts := new(Options)

	opts.HTTPMetricName = "test_metric"

	assert.Equal(t, "test_metric", opts.getHTTPMetricName())
}

func TestOptionsSetMetricsPath(t *testing.T) {
	opts := new(Options)

	opts.MetricsPath = "/test"

	assert.Equal(t, "/test", opts.getMetricsPath())
}

func TestOptionsLabels(t *testing.T) {
	opts := new(Options)

	opts.HTTPMetricDefaultLabelsNames = map[string]string{
		StatusCodeLabelName: "s",
		MethodLabelName:     "m",
	}

	opts.AdditionalHTTPMetricDefaultLabelsNames = map[string]string{"a": "a"}

	expectedLabels := []string{"a", "m", "path", "s"}

	sort.Strings(expectedLabels)

	for i, actual := range opts.getAllHTTPMetricDefaultLabelsNames() {
		assert.Equal(t, expectedLabels[i], actual)
	}
}

func TestOptionsSetHTTPMetricBuckets(t *testing.T) {
	opts := new(Options)

	expectedBuckets := []float64{0.1, 0.2, 0.3}
	opts.HTTPMetricBuckets = expectedBuckets

	for i, actual := range opts.getHTTPMetricBuckets() {
		assert.Equal(t, expectedBuckets[i], actual)
	}
}
