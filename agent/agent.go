package main

import (
	"log"
	"net/http"

	"github.com/pisxx/monitor/agent/web"
)

// // MetricsTempl provides template for metrics
// const metricsTempl = `{{ .Count }} metrics:
// {{ range .Metrics }}{
// 	"Metric": "{{ .Name }}",
// 	"Value": "{{ .Value }}"
// },
// {{ end }}`

func main() {
	http.HandleFunc("/", web.Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
