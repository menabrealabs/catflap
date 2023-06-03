// Catflapd is the execution service that runs with system level priviledges.
// It executes nftable firewall rules.
package main

import (
	"fmt"

	"github.com/menabrealabs/catflap/internal/pkg/port"
	"github.com/menabrealabs/catflap/internal/pkg/user"
)

func main() {
	user := user.New("Nyk", "dapper")
	user.Ports.Add(port.TCP, 22)
	user.Ports.Add(port.TCP, 80)

	fmt.Printf("%s pass is %x with ports %v\n", user.Name, user.Pass, user.Ports)
}
