package metrics

import (
	"runtime"
)

func GetCPUs() int {
	cpus := runtime.NumCPU()
	return cpus
}
