package port

// A TCP/IP port number: two byte integer
type Port uint16

// A finite set of ports
type Set map[Port]struct{}

// Add a Port to a port.Set
func (set Set) Add(port Port) {
	set[port] = struct{}{}
}

// Remove a Port from a port.Set
func (set Set) Remove(port Port) {
	delete(set, port)
}

// Clear all Ports from a port.Set
func (set Set) Reset() {
	set = Set{}
}

// Test whether a Port is in a port.Set
func (set Set) Contains(port Port) bool {
	_, exists := set[port]
	return exists
}
