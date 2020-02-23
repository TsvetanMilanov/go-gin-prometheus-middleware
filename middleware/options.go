package middleware

import (
	"sort"

	"github.com/prometheus/client_golang/prometheus"
)

// Options the options which will be used to create the middleware.
type Options struct {
	HTTPMetricName                  string
	HTTPMetricBuckets               []float64
	HTTPMetricUseNotNormalizedPaths bool
	MetricsPath                     string

	AdditionalHTTPMetricDefaultLabelsNames map[string]string
	HTTPMetricDefaultLabelsNames           map[string]string

	Registry    prometheus.Registerer
	Blacklister Blacklister
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
	result, ok := o.HTTPMetricDefaultLabelsNames[labelName]
	if !ok {
		return defaultHTTPMetricDefaultLabelsNamesNames[labelName]
	}

	return result
}

func (o *Options) getHTTPMetricDefaultLabelNames() []string {
	sorted := getSortedKeys(defaultHTTPMetricDefaultLabelsNamesNames)
	result := make([]string, len(sorted))
	for i, labelName := range sorted {
		result[i] = o.getLabelName(labelName)
	}

	sort.Strings(result)

	return result
}

func (o *Options) getAllHTTPMetricDefaultLabelsNames() []string {
	labels := append(o.getHTTPMetricDefaultLabelNames(), getSortedKeys(o.AdditionalHTTPMetricDefaultLabelsNames)...)

	sort.Strings(labels)

	return labels
}

func (o *Options) getRegistry() prometheus.Registerer {
	if o.Registry == nil {
		return prometheus.DefaultRegisterer
	}

	return o.Registry
}

func (o *Options) getBlacklister() Blacklister {
	return o.Blacklister
}
