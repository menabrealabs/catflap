// A port is an Internet Protocol (IP) port and can either
// be a TCP or UDP port. It is the sum of the protocol (TCP/UDP)
// and a two byte unsigned integer (the port number).
package port

// Protocol is an IP protocol (TCP or UDP).
type Protocol uint8

const (
	TCP Protocol = iota
	UDP
)

// A Port is a combination of IP protocol and IP port number.
type Port struct {
	Proto  Protocol // IP protocol (TCP or UDP).
	Number uint16   // IP port number.
}
