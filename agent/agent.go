package main

import (
	"log"
	"net/http"

	"github.com/pisxx/monitor/agent/register"

	"github.com/pisxx/monitor/agent/web"
)

func main() {
	reg, err := register.Agent()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(reg)
	http.HandleFunc("/", web.IndexMetrics)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
