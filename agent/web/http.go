package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pisxx/monitor/agent/metrics"
)

// Index - handler for index page
func Index(w http.ResponseWriter, r *http.Request) {
	metrics, err := metrics.GetMetrics()
	if err != nil {
		log.Fatal("Something went wrong")
	}
	// fmt.Println(metrics.Count)
	// var report = template.Must(template.New("metricsReport").Parse(metricsTempl))
	// report.Execute(os.Stdout, metrics)
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", " ")
	enc.Encode(metrics)
	fmt.Fprint(w, string(buf.Bytes()))
	// fmt.Println(string(buf.Bytes()))
}
