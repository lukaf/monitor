package monitor

import "testing"

var ds *DiskStat = &DiskStat{
	blockSize:  1024,
	blockTotal: 10,
	blockFree:  5,
}

func TestNewDiskStat(t *testing.T) {
	// This must NOT fail!
	var _ *DiskStat = NewDiskStat("")
}

func TestFree(t *testing.T) {
	free := ds.Free()
	if free != ds.blockFree*ds.blockSize {
		t.Errorf("Free value wrong: got %d, expected %d\n", free, ds.blockFree*ds.blockSize)
	}
}

func TestTotal(t *testing.T) {
	total := ds.Total()
	if total != ds.blockTotal*ds.blockSize {
		t.Errorf("Total value wrong: got %d, expected %d\n", total, ds.blockTotal*ds.blockSize)
	}
}

func TestUsed(t *testing.T) {
	used := ds.Used()
	if used != (ds.blockTotal-ds.blockFree)*ds.blockSize {
		t.Errorf("Used value wrong: got %d, expected %d\n", used, (ds.blockTotal-ds.blockFree)*ds.blockSize)
	}
}

func TestUsedPercent(t *testing.T) {
	percent := ds.UsedPercent()
	if percent != ((ds.blockTotal-ds.blockFree)*ds.blockSize)*100/(ds.blockTotal*ds.blockSize) {
		t.Errorf("UsedPercent value wrong: got %d, expected %d\n", percent, ((ds.blockTotal-ds.blockFree)*ds.blockSize)*100/(ds.blockTotal*ds.blockSize))
	}
}
