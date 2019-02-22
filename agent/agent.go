package main

import (
	"fmt"

	"github.com/pisxx/monitor/agent/metrics"
)

func main() {
	hostname := metrics.GetHostname()
	hostOS := metrics.GetOS()
	cpus := metrics.GetCPUs()
	mem := metrics.GetMemory()
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("Host OS: %s\n", hostOS)
	fmt.Printf("Host CPUs: %d\n", cpus)
	fmt.Printf("Host Memory: %d MB\n", mem/1024)
	// metrics.GetMemory()
}
