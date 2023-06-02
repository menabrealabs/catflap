package port

type Protocol uint8

const (
	TCP Protocol = iota
	UDP
)

// A TCP/IP port number: two byte integer
type Port struct {
	Proto  Protocol
	Number uint16
}

// A finite set of ports
type Set map[Port]struct{}

// Add a Port to a port.Set
func (set Set) Add(proto Protocol, number uint16) {
	port := Port{Proto: proto, Number: number}
	set[port] = struct{}{}
}

// Remove a Port from a port.Set
func (set Set) Remove(port Port) {
	delete(set, port)
}

// Clear all Ports from a port.Set
func (set Set) Reset() {
	for port := range set {
		delete(set, port)
	}
}

// Test whether a Port is in a port.Set
func (set Set) Contains(port Port) bool {
	_, exists := set[port]
	return exists
}
