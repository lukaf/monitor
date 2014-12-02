package main

import "strings"

var resources = map[string]func() (uint64, string, string){
	"uptime":          uptime,
	"memTotal":        memTotal,
	"memFree":         memFree,
	"memUsed":         memUsed,
	"memUsedPercent":  memUsedPercent,
	"procs":           procs,
	"load":            load,
	"diskFree":        diskFree,
	"diskTotal":       diskTotal,
	"diskUsed":        diskUsed,
	"diskUsedPercent": diskUsedPercent,
}

func uptime() (uint64, string, string) {
	return NewSysinfo().Uptime(), "Seconds", "uptime"
}

func memTotal() (uint64, string, string) {
	return NewSysinfo().MemTotal(), "Bytes", "totalMemory"
}

func memFree() (uint64, string, string) {
	return NewSysinfo().MemFree(), "Bytes", "freeMemory"
}

func memUsed() (uint64, string, string) {
	return NewSysinfo().MemUsed(), "Bytes", "usedMemory"
}

func memUsedPercent() (uint64, string, string) {
	return NewSysinfo().MemUsedPercent(), "Percent", "usedMemoryPercent"
}

func procs() (uint64, string, string) {
	return NewSysinfo().Procs(), "Count", "processes"
}

func load() (uint64, string, string) {
	return NewSysinfo().Load(), "Percent", "load"
}

func diskFree() (uint64, string, string) {
	return NewDiskStat("").Free(), "Bytes", "freeDisk"
}

func diskTotal() (uint64, string, string) {
	return NewDiskStat("").Total(), "Bytes", "totalDisk"
}

func diskUsed() (uint64, string, string) {
	return NewDiskStat("").Used(), "Bytes", "usedDisk"
}

func diskUsedPercent() (uint64, string, string) {
	return NewDiskStat("").UsedPercent(), "Percent", "usedDiskPercent"
}

func getResource(resource string) (uint64, string, string) {
	return resources[resource]()
}

func listResources() string {
	var resourceList []string
	for resource := range resources {
		resourceList = append(resourceList, resource)
	}

	return strings.Join(resourceList, ", ")
}
