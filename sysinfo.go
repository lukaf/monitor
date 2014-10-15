// +linux

package main

import (
	"log"
	"syscall"
)

type Sysinfo struct {
	uptime    uint64
	memTotal  uint64
	memFree   uint64
	memShared uint64
	memBuffer uint64
	swapTotal uint64
	swapFree  uint64
	load      [3]uint64
	procs     uint64
}

// Update updates the Sysinfo struct with current system information.
func (s *Sysinfo) Update() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Printf("Sysinfo call failed: %s", err.Error())
		return
	}

	s.uptime = uint64(info.Uptime)
	s.memTotal = info.Totalram
	s.memFree = info.Freeram
	s.memShared = info.Sharedram
	s.memBuffer = info.Bufferram
	s.swapTotal = info.Totalswap
	s.swapFree = info.Freeswap
	s.load = info.Loads
	s.procs = uint64(info.Procs)
}

// Uptime returns system uptime.
func (s Sysinfo) Uptime() uint64 {
	return s.uptime
}

// MemTotal returns all available memory.
func (s Sysinfo) MemTotal() uint64 {
	return s.memTotal
}

// MemFree returns amount of free memory in bytes.
func (s Sysinfo) MemFree() uint64 {
	return s.memFree + s.memBuffer
}

// MemUsed returns amount of used memory in bytes.
func (s Sysinfo) MemUsed() uint64 {
	return s.MemTotal() - s.MemFree()
}

// MemUsedPercent returns used memory in percent.
func (s Sysinfo) MemUsedPercent() uint64 {
	return s.MemUsed() * 100 / s.MemTotal()
}

// Procs returns number of currently running processes.
func (s Sysinfo) Procs() uint64 {
	return s.procs
}

// Load returns only 5 minute load average.
func (s Sysinfo) Load() uint64 {
	return s.load[1]
}

// NewSysinfo returns populated Sysinfo struct.
func NewSysinfo() *Sysinfo {
	s := &Sysinfo{}
	s.Update()

	return s
}
