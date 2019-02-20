package metrics

import (
	"log"
	"os"
)

// GetHostname returns name of host on which agetn is running
func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}
