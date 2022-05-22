package lib

import (
	"fmt"
	"testing"
)

func TestGetSystemConnection(t *testing.T) {
	_, err := GetSystemConnection()
	if err != nil {
		t.Error()
	}
}

func TestGetMemoryInfo(t *testing.T) {
	conn, _ := GetSystemConnection()
	dom, _ := conn.LookupDomainByName("vm1")
	meminfo, _ := GetGuestMemoryInfo(dom)
	fmt.Println(meminfo)
}

func TestAutoBalloon(t *testing.T) {
	conn, _ := GetSystemConnection()
	dom, _ := conn.LookupDomainByName("vm1")
	AutoBalloon(dom)
}
