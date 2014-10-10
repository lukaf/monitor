// +linux

package main

import (
	"log"
	"syscall"
)

const defaultPath = "/"

type DiskStat struct {
	blockSize  uint64
	blockTotal uint64
	blockFree  uint64
	path       string
}

func (ds *DiskStat) Update() {
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(ds.path, &stat)
	if err != nil {
		log.Printf("Statfs call failed: %s", err.Error())
		return
	}

	ds.blockSize = uint64(stat.Bsize)
	ds.blockTotal = stat.Blocks
	// Available (Bavail or f_bavail) instead of free (Bfree or f_bfree) as free
	// includes blocks reserved for the system and those cannot be used by unpriviledged
	// users.
	ds.blockFree = stat.Bavail
}

func (d DiskStat) Free() uint64 {
	return d.blockFree * d.blockSize
}

func (d DiskStat) Total() uint64 {
	return d.blockTotal * d.blockSize
}

func (d DiskStat) Used() uint64 {
	return d.Total() - d.Free()
}

func (d DiskStat) UsedPercent() uint64 {
	return d.Used() * 100 / d.Total()
}

func NewDiskStat(path string) *DiskStat {
	ds := &DiskStat{}
	if path == "" {
		ds.path = defaultPath
	}

	ds.Update()
	return ds
}
