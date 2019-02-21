package metrics

import (
	"fmt"
	"runtime"
)

func GetMemory() {
	// memory := runtime.M
	hostOS := runtime.GOOS
	switch hostOS {
	case "darwin":
		fmt.Printf("Operating system: %s\n", hostOS)
	case "linux":
		fmt.Printf("Operating system: %s\n", hostOS)
	}
}
