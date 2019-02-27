package metrics

import (
	"runtime"
	"strconv"
)

// GetCPUs returns cpu count
func GetCPUs() string {
	cpus := runtime.NumCPU()
	return strconv.Itoa(cpus)
}
