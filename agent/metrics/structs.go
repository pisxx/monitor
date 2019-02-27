package metrics

// Metrics struct
type MetricsStruct struct {
	Count   int             `json:"count"`
	Metrics []*MetricStruct `json:"metrics"`
}

// Metric struct
type MetricStruct struct {
	Name  string
	Value string
}
