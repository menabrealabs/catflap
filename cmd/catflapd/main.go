// Catflapd is the execution service that runs with system level priviledges.
// It executes nftable firewall rules.
package main

import (
	"fmt"
	"log"

	"github.com/menabrealabs/catflap/internal/pkg/port"
	"github.com/menabrealabs/catflap/internal/pkg/user"
)

func main() {
	user, err := user.New("Nyk", "dapper")
	if err != nil {
		log.Fatalf("failed to create user with error: %s\n", err)
	}

	user.Ports.Add(port.TCP, 22)
	user.Ports.Add(port.TCP, 80)

	fmt.Printf("%s pass is %x with ports %v\n", user.Name, user.Pass, user.Ports)
}
