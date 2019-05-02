// Poller Service
// This service queries SQS queue to get list of agents
// Based on the returned message it will try to access agent http endpoint to get metrics
// Service will use agent hostname field from SQS message
// It listens on port 9002
// To trigger polling you need to send GET to http://<ip>:9002

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pisxx/monitor/services/utils"
)

const (
	chooserQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/chooser-queue"
	// metricsQURL = "https://sqs.us-east-1.amazonaws.com/953738548419/metrics-queue"
)

func main() {

	ip := "0.0.0.0"
	port := "9002"
	http.HandleFunc("/", poll)
	log.Printf("Listening in %s:%s", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))

}

func poll(w http.ResponseWriter, req *http.Request) {

	// for {
	log.Print("Getting list of Agents from SQS\n")
	agents, messageID := utils.SQSGetAgentsList(chooserQURL)
	if *messageID == "Received no messages" {
		fmt.Fprintf(w, "Received no polling requests")
		return
	}
	// agents = []string{"onet.pl:10808"}
	fmt.Fprintf(w, "Polling metrics from: %v\n", agents)
	// fmt.Println(*messageID)

	metrics, _ := utils.PollMetrics(agents)
	// if err != nil {
	// 	fmt.Fprint(w, err.Error())
	// 	return
	// }
	fmt.Fprint(w, metrics)
	err := utils.DeleteMessage(messageID, chooserQURL)
	if err != nil {
		fmt.Fprintf(w, "Unable to poll  %s", *messageID)
	}
	// time.Sleep(5 * time.Second)
	// }
}
