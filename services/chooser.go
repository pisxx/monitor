package main

import (
	"log"
	"time"

	"github.com/pisxx/monitor/services/utils"
)

const (
	qURL = "https://sqs.us-east-1.amazonaws.com/953738548419/chooser-queue"
)

func main() {
	go forever()
	select {}
}

func forever() {
	for {
		agents := utils.GetAgents("agents_hostname")
		log.Printf("List of agents to send %s", agents)
		sendMessage := utils.SendAgentsList(agents, qURL)
		log.Printf("Message sent: %s", sendMessage)
		time.Sleep(5 * time.Second)
		fmt.Errorf(format string, a ...interface{})
	}
}
