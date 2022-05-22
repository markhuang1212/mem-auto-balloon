package lib

import (
	"fmt"
	"testing"
)

func TestGetSystemConnection(t *testing.T) {
	conn := GetSystemConnection()
	_, err := conn.GetHostname()
	if err != nil {
		t.Error()
	}
}

func TestGetMemoryInfo(t *testing.T) {
	conn := GetSystemConnection()
	dom, _ := conn.LookupDomainByName("vm1")
	meminfo := GetGuestMemoryInfo(dom)
	fmt.Println(meminfo)
}

func TestAutoBalloon(t *testing.T) {
	conn := GetSystemConnection()
	dom, _ := conn.LookupDomainByName("vm1")
	AutoBalloon(dom)
}
