package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// CreateMiddleware returns new gin middleware.
func CreateMiddleware(options *Options) gin.HandlerFunc {
	httpHistogramVec := createHistogram(options)

	return func(c *gin.Context) {
		if c.Request.URL.Path == options.getMetricsPath() {
			c.Next()
			return
		}

		start := time.Now()

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)

		labels := getHTTPMetricLabels(c, status, options)

		httpHistogramVec.With(labels).Observe(elapsed)
	}
}

func createHistogram(options *Options) *prometheus.HistogramVec {
	histogramOptions := prometheus.HistogramOpts{
		Name:    options.getHTTPMetricName(),
		Help:    "The temperature of the frog pond.",
		Buckets: options.getHTTPMetricBuckets(),
	}

	httpHistogramVec := prometheus.NewHistogramVec(histogramOptions, options.getAllHTTPMetricLabels())

	return httpHistogramVec
}

func getHTTPMetricLabels(c *gin.Context, status string, options *Options) map[string]string {
	result := make(map[string]string)

	result[StatusCodeLabelName] = status
	result[MethodLabelName] = c.Request.Method
	result[PathLabelName] = c.FullPath()

	for label, value := range options.CustomHTTPMetricLabels {
		result[label] = value
	}

	return result
}
