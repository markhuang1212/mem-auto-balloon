package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/markhuang1212/memdeflate/lib"
)

var ErrCommand = errors.New("bad command argument")

func main() {
	vmname := flag.String("vmname", "", "Name of the vm for auto ballooning")

	flag.Parse()

	if *vmname == "" {
		lib.FatalError(ErrCommand)
	}

	conn, err := lib.GetSystemConnection()
	if err != nil {
		lib.FatalError(err)
		return
	}

	dom, err := conn.LookupDomainByName(*vmname)

	if err != nil {
		lib.FatalError(lib.ErrNoDomain)
		return
	}

	err = lib.AutoBalloon(dom)

	if err != nil {
		lib.FatalError(err)
		return
	}

	fmt.Println("Success!")
}
