package job

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func MonitorResources() {
	log.Println("=====START MONITOR RESOURCES=====")
	// Get CPU usage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error getting CPU usage:", err)
	} else {
		fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
	}

	// Get memory usage
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory usage:", err)
	} else {
		fmt.Printf("Memory Usage: %.2f%%\n", vmStat.UsedPercent)
	}

	// Get number of goroutines
	fmt.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())
	log.Println("=====END MONITOR RESOURCES=====")
}
