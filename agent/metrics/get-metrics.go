package metrics

// GetMetrics getts all metrics from metric package
func GetMetrics() (*MetricsStruct, error) {
	// Add functions to metrics dictionary
	type fn func() string
	m := make(map[string]fn)
	m["Hostname"] = GetHostname
	m["HostOS"] = GetOS
	m["CPU"] = GetCPUs
	m["Memory"] = GetMemory
	var result MetricsStruct
	result.Count = len(m)
	// var metricsSlice []web.MetricStructStr
	for k, v := range m {
		var metric MetricStruct
		metric.Name = k
		metric.Value = v()
		result.Metrics = append(result.Metrics, &metric)
	}
	return &result, nil
}
