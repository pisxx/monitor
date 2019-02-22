package metrics

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// GetMemory get total memory from guest using correct functions depending on which is it running
func GetMemory() int {
	var memSize int
	hostOS := runtime.GOOS
	switch hostOS {
	case "darwin":
		memSize = darwinGetMemory()
	case "linux":
		// fmt.Printf("Operating system: %s\n", hostOS)
		memSize = linuxGetMemory()
		// memSize = 1
	}
	return memSize
}

// func darwinGetMemory() {
// 	getMem := exec.Command("system_profiler", "SPHardwareDataType")
// 	parseMem1 := exec.Command("grep", "Memory")
// 	r, w := io.Pipe()
// 	getMem.Stdout = w
// 	parseMem1.Stdin = r

// 	var b2 bytes.Buffer
// 	parseMem1.Stdout = &b2

// 	getMem.Start()
// 	parseMem1.Start()
// 	getMem.Wait()
// 	w.Close()
// 	parseMem1.Wait()
// 	io.Copy(os.Stdout, &b2)
// }

func darwinGetMemory() int {
	out, err := exec.Command("sysctl", "-n", "hw.memsize").Output()
	if err != nil {
		log.Fatal(err)
	}
	mem, err := strconv.ParseUint(strings.TrimSpace(string(out)), 10, 64)
	memKB := int(mem / 1024)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Type of memSize: %T\n", memKB)
	// fmt.Printf("Memory: %d\n", memKB)
	return memKB

}

func linuxGetMemory() int {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	memTotal := scanner.Text()
	memTotalSplitted := strings.Split(memTotal, ":")
	memSize := strings.TrimLeft(memTotalSplitted[1], " ")
	memSizeSplitted := strings.Split(memSize, " ")
	memSizeInt, err := strconv.ParseUint(memSizeSplitted[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(memSizeInt)
}
