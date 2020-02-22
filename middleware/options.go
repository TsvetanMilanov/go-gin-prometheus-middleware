package middleware

import "sort"

// Options the options which will be used to create the middleware.
type Options struct {
	HTTPMetricName    string
	HTTPMetricBuckets []float64
	MetricsPath       string

	CustomHTTPMetricLabels   map[string]string
	HTTPMetricLabelsOverride map[string]string
}

func (o *Options) getHTTPMetricName() string {
	if len(o.HTTPMetricName) == 0 {
		return defaultHTTPMetricName
	}

	return o.HTTPMetricName
}

func (o *Options) getMetricsPath() string {
	if len(o.MetricsPath) == 0 {
		return defaultMetricsPath
	}

	return o.MetricsPath
}

func (o *Options) getHTTPMetricBuckets() []float64 {
	if len(o.HTTPMetricBuckets) == 0 {
		return defaultHTTPMetricBuckets
	}

	return o.HTTPMetricBuckets
}

func (o *Options) getLabelName(labelName string) string {
	result, ok := o.HTTPMetricLabelsOverride[labelName]
	if !ok {
		return defaultHTTPMetricLabelsNames[labelName]
	}

	return result
}

func (o *Options) getHTTPMetricDefaultLabelNames() []string {
	sorted := getSortedKeys(defaultHTTPMetricLabelsNames)
	result := make([]string, len(sorted))
	for i, labelName := range sorted {
		result[i] = o.getLabelName(labelName)
	}

	sort.Strings(result)

	return result
}

func (o *Options) getAllHTTPMetricLabels() []string {
	labels := append(o.getHTTPMetricDefaultLabelNames(), getSortedKeys(o.CustomHTTPMetricLabels)...)

	sort.Strings(labels)

	return labels
}
