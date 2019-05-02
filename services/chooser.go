// Chooser service
// This service queries DynamoDB to get registered agents
// It will then put SQS message with agents for the Poller service to consume
// It listens on port 9000
// To trigger choosing you need to send GET to http://<ip>:9000
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pisxx/monitor/services/utils"
)

const (
	qURL = "https://sqs.us-east-1.amazonaws.com/953738548419/chooser-queue"
)

func main() {
	ip := "0.0.0.0"
	port := "9000"
	http.HandleFunc("/", choose)
	log.Printf("Listening in %s:%s", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))
	// go forever()
	// select {}
}

func choose(w http.ResponseWriter, req *http.Request) {
	// for {
	agents := utils.DBGetAgents("agents_hostname", -1)
	log.Print("Choosed agents: ")
	fmt.Fprintf(w, "Choosed agents:\n")
	for _, agent := range agents {
		fmt.Fprintf(w, "%s\n", agent)
	}
	sendMessage := utils.SQSSendAgentsList(agents, qURL)
	log.Printf("Message sent: %s", sendMessage)
	fmt.Fprintf(w, "Message sent: %v", sendMessage)
	time.Sleep(5 * time.Second)
	// fmt.Errorf(format string, a ...interface{})
	// }
}
