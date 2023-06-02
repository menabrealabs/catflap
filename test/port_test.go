package port_test

import (
	"testing"

	"github.com/menabrealabs/catflap/internal/pkg/port"
)

var (
	tcpCase = port.Port{Proto: port.TCP, Number: 22}
	udpCase = port.Port{Proto: port.UDP, Number: 25}
)

func setupPorts() port.Set {
	ports := port.Set{}
	ports.Add(port.TCP, 22)
	ports.Add(port.UDP, 25)

	return ports
}

func TestPortAdd(t *testing.T) {
	ports := setupPorts()

	t.Run("should add TCP port to set", func(t *testing.T) {
		_, found := ports[tcpCase]
		if !found {
			t.Fail()
		}
	})

	t.Run("should add UDP port to set", func(t *testing.T) {
		_, found := ports[udpCase]
		if !found {
			t.Fail()
		}
	})
}

func TestPortRemove(t *testing.T) {
	ports := setupPorts()
	t.Run("should delete TCP port from set", func(t *testing.T) {
		ports.Remove(tcpCase)

		_, found := ports[tcpCase]
		if found {
			t.Error("port did not get removed from the set")
		}
	})

	t.Run("should delete UDP port from set", func(t *testing.T) {
		port := udpCase
		ports.Remove(port)

		_, found := ports[port]
		if found {
			t.Error("port did not get removed from the set")
		}
	})

	t.Run("should not delete other ports from set", func(t *testing.T) {
		ports := setupPorts()
		ports.Remove(tcpCase)

		_, found := ports[udpCase]
		if !found {
			t.Error("port removed from set when other port removed")
		}
	})
}

func TestPortContains(t *testing.T) {
	ports := setupPorts()
	t.Run("should return true when port is in the set", func(t *testing.T) {
		ports.Add(port.TCP, 22)

		if ports.Contains(tcpCase) == false {
			t.Error("failed to report that a port exists in the set")
		}
	})

	t.Run("should return false when port is not in the set", func(t *testing.T) {
		ports.Remove(tcpCase)

		if ports.Contains(tcpCase) == true {
			t.Error("failed to report that a port exists in the set")
		}
	})
}
