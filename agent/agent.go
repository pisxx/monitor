package main

import (
	"fmt"

	"github.com/pisxx/monitor/agent/metrics"
)

func main() {
	hostname := metrics.GetHostname()
	cpus := metrics.GetCPUs()
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("Number of CPUs: %d\n", cpus)
	metrics.GetMemory()
}
