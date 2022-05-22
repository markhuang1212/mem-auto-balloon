package lib

import (
	"errors"
	"fmt"
	"os"
	"time"

	"libvirt.org/go/libvirt"
)

var ErrConnection = errors.New("cannot connect to libvirt")
var ErrGetMemInfo = errors.New("cannot get guest memory statistics")
var ErrBadMemInfo = errors.New("bad memory statistics")
var ErrNotRunning = errors.New("vm not running")
var ErrNoDomain = errors.New("no domain with given name")

func FatalError(e error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.Error())
	os.Exit(1)
}

func GetSystemConnection() (*libvirt.Connect, error) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return nil, ErrConnection
	}
	return conn, nil
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

func GetGuestMemoryInfo(dom *libvirt.Domain) (MemoryInfo, error) {

	stats, err := dom.MemoryStats(16, 0)

	if err != nil {
		return MemoryInfo{}, ErrGetMemInfo
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
			return MemoryInfo{}, ErrBadMemInfo
		}
	}

	return ret, nil
}

func AutoBalloon(dom *libvirt.Domain) error {

	state, _, err := dom.GetState()
	if err != nil {
		return err
	}

	if state != libvirt.DOMAIN_RUNNING {
		return ErrNotRunning
	}

	maxmem, err := dom.GetMaxMemory()
	if err != nil {
		return err
	}

	meminfo, err := GetGuestMemoryInfo(dom)
	if err != nil {
		return err
	}

	target := meminfo.ActualBalloon - meminfo.Usable

	fmt.Printf("Setting Memory From %dM to %dM\n", meminfo.ActualBalloon/1024, target/1024)

	dom.SetMemory(meminfo.ActualBalloon - meminfo.Usable)

	time.Sleep(5 * time.Second)

	fmt.Printf("Setting Memory From %dM to %dM\n", target/1024, maxmem/1024)

	dom.SetMemory(maxmem)

	return nil
}
