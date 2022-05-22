package main

import (
	"flag"
	"fmt"

	"github.com/markhuang1212/memdeflate/lib"
)

func main() {
	vmname := flag.String("vmname", "", "Name of the vm for auto ballooning")

	flag.Parse()

	if *vmname == "" {
		panic("No VM name")
	}

	conn := lib.GetSystemConnection()
	dom, err := conn.LookupDomainByName(*vmname)

	if err != nil {
		panic("Cannot find domain")
	}

	lib.AutoBalloon(dom)

	fmt.Println("Success!")
}
