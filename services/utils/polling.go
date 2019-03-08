package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PollMetrics(qURL string, agents []string) {
	// Poll from agents
	// var metricsSlice []utils.MetricsStruct
	fmt.Println()
	// type M map[string]string
	metricsMapList := make([]map[string]string, 3)
	// var metricsMap M
	metricsMap := make(map[string]string)
	for _, i := range agents {
		var metrics MetricsStruct
		fmt.Printf("Polling metrics from %s\n", i)
		resp, err := http.Get("http://" + i)
		if err != nil {
			log.Fatal("Unable to poll metrics")
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&metrics)
		if err != nil {
			log.Fatal("issue")
		}
		for _, v := range metrics.Metrics {
			fmt.Printf("%s: %s\t", v.Name, v.Value)
			metricsMap[v.Name] = v.Value
			metricsMapList = append(metricsMapList, metricsMap)
		}
		fmt.Println()
		fmt.Println()
	}
	// fmt.Println(message)
	fmt.Println("Metrics polled, deleting poll request from SQS")
	// return metricsMapList
}
