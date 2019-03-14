package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pisxx/monitor/services/utils"
)

const (
	chooserQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/chooser-queue"
	// metricsQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/metrics-queue"
)

func main() {

	// go poll()
	// select {}
	http.HandleFunc("/", poll)
	log.Print("Listening on 0.0.0.0:9000")
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))

}

func poll(w http.ResponseWriter, req *http.Request) {
	for {
		// fmt.Fprintf("Getting list of Agents from SQS\n")
		listOfAgnets, messageID := utils.GetAgentsList(chooserQURL)
		if *messageID == "Received no messages" {
			fmt.Fprintf(w, "Received no messages")
			return
		}
		fmt.Fprintf(w, "List of agents to poll metrics from: %v\n", listOfAgnets)
		// fmt.Println(*messageID)
		err := utils.DeleteMessage(messageID, chooserQURL)
		if err != nil {
			fmt.Fprintf(w, "Unable to delete message %s", *messageID)
		}
		utils.PollMetrics(chooserQURL, listOfAgnets)
		time.Sleep(5 * time.Second)
	}
}
