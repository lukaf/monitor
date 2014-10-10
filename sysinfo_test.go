package main

import "testing"

var s *Sysinfo = &Sysinfo{
	uptime:    100,
	memTotal:  100,
	memFree:   50,
	memShared: 20,
	memBuffer: 10,
	swapTotal: 5,
	swapFree:  2,
	load:      [3]uint64{1, 2, 3},
}

func TestNewSysinfo(t *testing.T) {
	// This must NOT fail!
	var _ *Sysinfo = NewSysinfo()
}

func TestUptime(t *testing.T) {
	uptime := s.Uptime()
	if uptime != s.uptime {
		t.Errorf("Uptime value wrong: got %d, expected %d\n", uptime, s.uptime)
	}
}

func TestMemTotal(t *testing.T) {
	mem := s.MemTotal()
	if mem != s.memTotal {
		t.Errorf("MemTotal value wrong: got %d, expected %d\n", mem, s.memTotal)
	}
}

func TestMemFree(t *testing.T) {
	mem := s.MemFree()
	if mem != (s.memFree + s.memBuffer) {
		t.Errorf("MemFree value wrong: got %d, expected %d\n", mem, s.memFree+s.memBuffer)
	}
}

func TestMemUsed(t *testing.T) {
	mem := s.MemUsed()
	if mem != (s.memTotal - (s.memFree + s.memBuffer)) {
		t.Errorf("MemUsed value wrong: got %d, expected %d", mem, s.memTotal-(s.memFree+s.memBuffer))
	}
}

func TestMemUsedPercent(t *testing.T) {
	mem := s.MemUsedPercent()
	if mem != ((s.memTotal - (s.memFree + s.memBuffer)) * 100 / s.memTotal) {
		t.Errorf("MemUsedPercent value wrong: got %d, expected %d\n", mem, (s.memTotal-(s.memFree+s.memBuffer))*100/s.memTotal)
	}
}
