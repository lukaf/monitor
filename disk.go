// +linux

package main

import (
	"log"
	"syscall"
)

const defaultPath = "/"

// The DiskStat type represents the current filesystem status.
type DiskStat struct {
	blockSize  uint64
	blockTotal uint64
	blockFree  uint64
	path       string
}

// Update updates the information about filesystem.
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

// Free returns the amount of filesystem in bytes available to unpriviledged user.
func (d DiskStat) Free() uint64 {
	return d.blockFree * d.blockSize
}

// Total returns the size of the filesystem in bytes.
func (d DiskStat) Total() uint64 {
	return d.blockTotal * d.blockSize
}

// Used returns the amount of filesystem used in bytes.
func (d DiskStat) Used() uint64 {
	return d.Total() - d.Free()
}

// UsedPercent returns the amount of filesystem used by percent.
func (d DiskStat) UsedPercent() uint64 {
	return d.Used() * 100 / d.Total()
}

// NewDiskStat returns a new populated DiskStat struct.
// If path is a non empty string DiskStat will hold the information about filesystem
// mounted under the mountpoint belonging to path.
func NewDiskStat(path string) *DiskStat {
	ds := &DiskStat{}
	if path == "" {
		ds.path = defaultPath
	}

	ds.Update()
	return ds
}
