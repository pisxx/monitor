package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type metricsMapFor map[string]string
type metricsMap map[string]interface{}
type metricsMapList []metricsMap

func PollMetrics(agents []string) (metricsMapList, error) {
	// Poll from agents
	// var metricsSlice []utils.MetricsStruct
	fmt.Println()
	// type M map[string]string
	metricsMapList := make([]metricsMap, 0)
	// var metricsMap M
	// metricsMap := make(map[string]string)
	// metricsMap := make(MetricsMap)
	metricsMapFor := make(metricsMapFor)
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	for _, i := range agents {
		var metrics MetricsStruct
		metricsMap := make(metricsMap)
		log.Printf("Polling metrics from %s\n", i)
		metricsMap["agent"] = i
		resp, err := client.Get("http://" + i)
		if err != nil {
			log.Print("Unable to poll metrics")
			log.Print(err)
			metricsMap["status"] = err.Error()
			metricsMapList = append(metricsMapList, metricsMap)
			continue
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&metrics)
		if err != nil {
			// log.Fatal("issue")
			return nil, err

		}
		for _, v := range metrics.Metrics {
			fmt.Printf("%s: %s\t", v.Name, v.Value)
			metricsMapFor[v.Name] = v.Value
			// metricsMapList = append(metricsMapList, metricsMap)
		}
		metricsMap["metrics"] = metricsMapFor
		fmt.Println()
		fmt.Println()
		// metricsMapList = append(metricsMapList, metricsMap)
		resp.Body.Close()
	}
	// fmt.Println(message)
	// log.Print("-------- After loop ---------")
	// log.Print("Length of MetricsMap -> ", len(metricsMapList))
	log.Print(metricsMapList)
	// fmt.Fprint(w, "Metrics polled, deleting poll request from SQS")
	// return metricsMapList
	return metricsMapList, nil
}
