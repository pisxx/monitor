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
	port := "9001"
	http.HandleFunc("/", choose)
	log.Printf("Listening in %s:%s", ip, port)
	log.Fatal(http.ListenAndServe(ip+":"+port, nil))
	// go forever()
	// select {}
}

func choose(w http.ResponseWriter, req *http.Request) {
	// for {
	agents := utils.DBGetAgents("agents_hostname")
	log.Printf("List of agents to send %s", agents)
	fmt.Fprintf(w, "List of agents to send %s\n", agents)
	sendMessage := utils.SQSSendAgentsList(agents, qURL)
	log.Printf("Message sent: %s", sendMessage)
	fmt.Fprintf(w, "Message sent: %v", sendMessage)
	time.Sleep(5 * time.Second)
	// fmt.Errorf(format string, a ...interface{})
	// }
}
