package main

import (
	"fmt"
	"log"

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
	fmt.Printf("Getting list of Agents from SQS\n")
	listOfAgnets, messageID := utils.GetAgentsList(chooserQURL)
	fmt.Printf("List of agents to poll metrics from: %v\n", listOfAgnets)
	// fmt.Println(*messageID)
	err := utils.DeleteMessage(messageID, chooserQURL)
	if err != nil {
		log.Printf("Unable to delete message %s", *messageID)
	}
	utils.PollMetrics(chooserQURL, listOfAgnets)
}
