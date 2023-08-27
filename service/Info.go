package info

import (
	"encoding/json"
	"runtime"
	"time"
)

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInfo struct {
	OS        string `json:"os"`
	OSVersion string `json:"os_version"`
	Timestamp int64  `json:"timestamp"`
	Memory    uint64 `json:"memory"`
	CPUInfo   string `json:"cpu_info"`
	DiskSpace uint64 `json:"disk_space"`
}

func GetSystemInfo() ([]byte, error) {
	osInfo := runtime.GOOS
	osVersion := ""
	timestamp := time.Now().Unix()

	// Get OS version for Unix-like systems
	if osInfo == "linux" || osInfo == "darwin" {
		//uname := syscall.Utsname{}
		//if err := syscall.Uname(&uname); err == nil {
		//	osVersion = string(uname.Release[:])
		//}
		osVersion = ""
	}

	memory, _ := mem.VirtualMemory()
	cpuInfo, _ := cpu.Info()
	diskSpace, _ := disk.Usage("/")

	info := SystemInfo{
		OS:        osInfo,
		OSVersion: osVersion,
		Timestamp: timestamp,
		Memory:    memory.Total,
		CPUInfo:   cpuInfo[0].String(),
		DiskSpace: diskSpace.Total,
	}

	return json.MarshalIndent(info, "", "  ")
}

func main() {
	systemInfo, err := GetSystemInfo()
	if err != nil {
		panic(err)
	}

	// Print the JSON formatted system information
	println(string(systemInfo))
}
