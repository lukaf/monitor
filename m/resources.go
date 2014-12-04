package monitor

import "strings"

var Resources = map[string]func() (uint64, string, string){
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
	return NewSysinfo().MemTotal(), "Bytes", "memoryTotal"
}

func memFree() (uint64, string, string) {
	return NewSysinfo().MemFree(), "Bytes", "memoryFree"
}

func memUsed() (uint64, string, string) {
	return NewSysinfo().MemUsed(), "Bytes", "memoryUsed"
}

func memUsedPercent() (uint64, string, string) {
	return NewSysinfo().MemUsedPercent(), "Percent", "memoryUsedPercent"
}

func procs() (uint64, string, string) {
	return NewSysinfo().Procs(), "Count", "processes"
}

func load() (uint64, string, string) {
	return NewSysinfo().Load(), "Percent", "load"
}

func diskFree() (uint64, string, string) {
	return NewDiskStat("").Free(), "Bytes", "diskFree"
}

func diskTotal() (uint64, string, string) {
	return NewDiskStat("").Total(), "Bytes", "diskTotal"
}

func diskUsed() (uint64, string, string) {
	return NewDiskStat("").Used(), "Bytes", "diskUsed"
}

func diskUsedPercent() (uint64, string, string) {
	return NewDiskStat("").UsedPercent(), "Percent", "diskUsedPercent"
}

func GetResource(resource string) (uint64, string, string) {
	return Resources[resource]()
}

func ListResources() string {
	var resourceList []string
	for resource := range Resources {
		resourceList = append(resourceList, resource)
	}

	return strings.Join(resourceList, ", ")
}
