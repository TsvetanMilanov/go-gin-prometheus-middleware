package middleware

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// New creates new gin middleware with the default options.
func New() gin.HandlerFunc {
	options := new(Options)

	handler := NewWithOptions(options)

	return handler
}

// NewWithOptions creates new gin middleware with the provided options.
func NewWithOptions(options *Options) gin.HandlerFunc {
	httpHistogramVec := createHistogram(options)

	options.getRegistry().MustRegister(httpHistogramVec)

	return func(c *gin.Context) {
		if c.Request.URL.Path == options.getMetricsPath() {
			c.Next()
			return
		}

		start := time.Now()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)

		labels := getHTTPMetricDefaultLabelsNames(c, status, options)

		httpHistogramVec.With(labels).Observe(elapsed)
	}
}

func createHistogram(options *Options) *prometheus.HistogramVec {
	histogramOptions := prometheus.HistogramOpts{
		Name:    options.getHTTPMetricName(),
		Help:    fmt.Sprintf("Duration summary of http responses labeled with: %s", strings.Join(options.getAllHTTPMetricDefaultLabelsNames(), ", ")),
		Buckets: options.getHTTPMetricBuckets(),
	}

	httpHistogramVec := prometheus.NewHistogramVec(histogramOptions, options.getAllHTTPMetricDefaultLabelsNames())

	return httpHistogramVec
}

func getHTTPMetricDefaultLabelsNames(c *gin.Context, status string, options *Options) map[string]string {
	result := make(map[string]string)

	result[StatusCodeLabelName] = status
	result[MethodLabelName] = c.Request.Method

	pathValue := c.FullPath()

	if options.HTTPMetricUseNotNormalizedPaths {
		pathValue = c.Request.URL.Path
	}

	result[PathLabelName] = pathValue

	for label, value := range options.AdditionalHTTPMetricDefaultLabelsNames {
		result[label] = value
	}

	return result
}
