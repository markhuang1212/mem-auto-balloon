package lib

import (
	"time"

	"libvirt.org/go/libvirt"
)

func GetSystemConnection() *libvirt.Connect {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic("Cannot create connection")
	}
	return conn
}

type MemoryInfo struct {
	SwapIn        uint64
	SwapOut       uint64
	MajorFault    uint64
	MinorFault    uint64
	Unused        uint64
	Available     uint64
	ActualBalloon uint64
	RSS           uint64
	Usable        uint64
	LastUpdate    uint64
	DiskCaches    uint64
}

func GetGuestMemoryInfo(dom *libvirt.Domain) MemoryInfo {
	stats, err := dom.MemoryStats(16, 0)

	if err != nil {
		panic("Cannot get MemoryStats")
	}

	ret := MemoryInfo{}

	valid := make([]bool, 11)
	for index := range valid {
		valid[index] = false
	}

	for _, stat := range stats {
		if stat.Tag < 11 {
			valid[stat.Tag] = true
		}
		switch stat.Tag {
		case 0:
			ret.SwapIn = stat.Val
		case 1:
			ret.SwapOut = stat.Val
		case 2:
			ret.MajorFault = stat.Val
		case 3:
			ret.MinorFault = stat.Val
		case 4:
			ret.Unused = stat.Val
		case 5:
			ret.Available = stat.Val
		case 6:
			ret.ActualBalloon = stat.Val
		case 7:
			ret.RSS = stat.Val
		case 8:
			ret.Usable = stat.Val
		case 9:
			ret.LastUpdate = stat.Val
		case 10:
			ret.DiskCaches = stat.Val
		}
	}

	for _, val := range valid {
		if !val {
			panic("Invalid Memory Information")
		}
	}

	return ret
}

func AutoBalloon(dom *libvirt.Domain) {
	maxmem, err := dom.GetMaxMemory()
	if err != nil {
		panic("Cannot get Guest Max Memory")
	}

	meminfo := GetGuestMemoryInfo(dom)

	dom.SetMemory(meminfo.ActualBalloon - meminfo.Usable)

	time.Sleep(5 * time.Second)

	dom.SetMemory(maxmem)
}
