package main

import (
	"fmt"

	"github.com/pisxx/monitor/services/utils"
)

const (
	chooserQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/chooser-queue"
	// metricsQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/metrics-queue"
)

func main() {

	go poll()
	select {}

}

func poll() {
	metrics := utils.GetAgentsList(chooserQURL)
	fmt.Println(metrics)
}
