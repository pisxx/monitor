package metrics

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func GetMemory() int {
	var memSize int
	hostOS := runtime.GOOS
	switch hostOS {
	case "darwin":
		memSize = darwinGetMemory()
	case "linux":
		fmt.Printf("Operating system: %s\n", hostOS)
		linuxGetMemory()
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

func linuxGetMemory() {
	getMem := exec.Command("cat", "/proc/info")
	parseMem1 := exec.Command("grep", "MemTotal:")
	r, w := io.Pipe()
	getMem.Stdout = w
	parseMem1.Stdin = r

	var b2 bytes.Buffer
	parseMem1.Stdout = &b2

	getMem.Start()
	parseMem1.Start()
	getMem.Wait()
	w.Close()
	parseMem1.Wait()
	io.Copy(os.Stdout, &b2)
}
