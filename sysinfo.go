// +linux

package main

import (
	"log"
	"syscall"
)

type Sysinfo struct {
	uptime    int64
	memTotal  uint64
	memFree   uint64
	memShared uint64
	memBuffer uint64
	swapTotal uint64
	swapFree  uint64
	load      [3]uint64
}

func (s *Sysinfo) Update() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Printf("Sysinfo call failed: %s", err.Error())
		return
	}

	s.uptime = info.Uptime
	s.memTotal = info.Totalram
	s.memFree = info.Freeram
	s.memShared = info.Sharedram
	s.memBuffer = info.Bufferram
	s.swapTotal = info.Totalswap
	s.swapFree = info.Freeswap
	s.load = info.Loads
}

func (s Sysinfo) Uptime() int64 {
	return s.uptime
}

func (s Sysinfo) MemTotal() uint64 {
	return s.memTotal
}

func (s Sysinfo) MemFree() uint64 {
	return s.memFree + s.memBuffer
}

func (s Sysinfo) MemUsed() uint64 {
	return s.MemTotal() - s.MemFree()
}

func (s Sysinfo) MemUsedPercent() uint64 {
	return s.MemUsed() * 100 / s.MemTotal()
}

func NewSysinfo() *Sysinfo {
	s := &Sysinfo{}
	s.Update()

	return s
}
