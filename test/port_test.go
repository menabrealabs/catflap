package port_test

import (
	"testing"

	"github.com/menabrealabs/catflap/internal/pkg/port"
)

func TestPortAdd_adds_tcp_port_to_set(t *testing.T) {
	ports := port.Set{}
	ports.Add(port.TCP, 22)
	ports.Add(port.UDP, 25)

	expected := port.Port{Proto: port.TCP, Number: 22}
	_, found := ports[expected]

	if !found {
		t.Error("port did not get added to the set")
	}
}

func TestPortAdd_adds_udp_port_to_set(t *testing.T) {
	ports := port.Set{}
	ports.Add(port.TCP, 22)
	ports.Add(port.UDP, 25)

	expected := port.Port{Proto: port.UDP, Number: 25}
	_, found := ports[expected]

	if !found {
		t.Error("port did not get added to the set")
	}
}
