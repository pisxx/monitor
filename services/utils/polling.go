package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type metricsMapForType map[string]string
type metricsType map[string]interface{}
type metricsMapType map[string]metricsType

// type agentDetails map[string]string
// type agentDetailsMap map[string]agentDetails

func PollMetrics(agents agentDetailsMap) (metricsMapType, error) {
	// Poll from agents
	// var metricsSlice []utils.MetricsStruct
	// fmt.Println()
	// type M map[string]string
	metricsMapVar := make(metricsMapType, 0)
	// var metricsMap M
	// metricsMap := make(map[string]string)
	// metricsMap := make(MetricsMap)
	metricsMapForVar := make(metricsMapForType)
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	for _, v := range agents {
		details := make(agentDetailsP, 0)
		var metricsStruct MetricsStruct
		details = *v
		metricsVar := make(metricsType, 0)
		// endpoint := details["hostname"]
		var endpoint string
		// log.Printf("----- Polling metrics from %s -----\n", endpoint)
		port := details["port"]
		if details["env"] == "docker" {
			endpoint = details["ip"]
			log.Printf("INFO: Polling metrics from docker: %s:%s\n", endpoint, port)
		} else if details["env"] == "aws-vm" {
			endpoint = details["public_ip"]
			// fmt.Print("------ This is AWS IP : ", details)
			log.Printf("INFO: Polling metrics from aws vm: %s:%s\n", endpoint, port)
		}
		resp, err := client.Get("http://" + endpoint + ":" + port)
		if err != nil {
			log.Print("ERROR: Unable to poll metrics")
			log.Print("ERROR: ", err)
			metricsVar["status"] = err.Error()
			// metricsMapList = append(metricsMapList, metricsMap)
			hostname := details["hostname"]
			metricsMapVar[hostname] = metricsVar
			continue
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&metricsStruct)
		if err != nil {
			// log.Fatal("issue")
			return nil, err

		}
		for _, v := range metricsStruct.Metrics {
			fmt.Printf("%s: %s\t", v.Name, v.Value)
			metricsMapForVar[v.Name] = v.Value
			// metricsMapList = append(metricsMapList, metricsMap)
		}
		metricsVar["metrics"] = metricsMapForVar
		fmt.Println()
		fmt.Println()
		metricsMapVar[details["hostname"]] = metricsVar
		// metricsMapList = append(metricsMapList, metricsMap)
		resp.Body.Close()
	}
	// fmt.Println(message)
	// log.Print("-------- After loop ---------")
	// log.Print("Length of MetricsMap -> ", len(metricsMapList))
	// log.Print(metricsMapList)
	// fmt.Fprint(w, "Metrics polled, deleting poll request from SQS")
	// return metricsMapList
	return metricsMapVar, nil
}
