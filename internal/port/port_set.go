package port

// A Set is a finite set of Port objects.
type Set map[Port]struct{}

// Add appends a Port object to a port.Set.
func (set Set) Add(proto Protocol, number uint16) {
	port := Port{Proto: proto, Number: number}
	set[port] = struct{}{}
}

// Remove deletes a Port object from a port.Set.
func (set Set) Remove(port Port) {
	delete(set, port)
}

// Clear deletes all Ports objects from a port.Set.
func (set Set) Reset() {
	for port := range set {
		delete(set, port)
	}
}

// Contains tests whether a Port object is an element of a port.Set.
func (set Set) Contains(port Port) bool {
	_, exists := set[port]
	return exists
}
