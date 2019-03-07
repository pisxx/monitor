package main

import (
	"log"

	aws "github.com/pisxx/monitor/services/utils"
)

const (
	qURL = "https://sqs.us-east-1.amazonaws.com/953738548419/puller-queue"
)

func main() {
	agents := utils.GetAgents("agents_hostname")
	log.Printf("List of agents to send %s", agents)
	sendMessage := utils.SendAgentsList(agents, qURL)
	log.Printf("Message sent: %s", sendMessage)
}
