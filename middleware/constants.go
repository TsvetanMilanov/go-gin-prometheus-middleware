package middleware

const (
	// StatusCodeLabelName The default label name for the request status code.
	StatusCodeLabelName = "status_code"
	// MethodLabelName The default label name for the request method.
	MethodLabelName = "method"
	// PathLabelName The default label name for the request path.
	PathLabelName = "path"

	defaultHTTPMetricName = "http_request_duration_seconds"
	defaultMetricsPath    = "/metrics"
)

var (
	defaultHTTPMetricBuckets     = []float64{0.1, 0.25, 0.5, 1, 2, 5, 10, 15, 20, 30}
	defaultHTTPMetricLabelsNames = map[string]string{
		StatusCodeLabelName: StatusCodeLabelName,
		MethodLabelName:     MethodLabelName,
		PathLabelName:       PathLabelName,
	}
)
