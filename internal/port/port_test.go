package port

import (
	"testing"
)

var (
	tcpCase = Port{Proto: TCP, Number: 22}
	udpCase = Port{Proto: UDP, Number: 25}
)

func setupPorts() Set {
	ports := Set{}
	ports.Add(TCP, 22)
	ports.Add(UDP, 25)

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
		ports.Add(TCP, 22)

		if ports.Contains(tcpCase) == false {
			t.Error("failed to report that a port exists in the set")
		}
	})

	t.Run("should return false when port is not in the set", func(t *testing.T) {
		ports.Remove(tcpCase)

		if ports.Contains(tcpCase) == true {
			t.Error("wrongly reports that a port exists in the set")
		}
	})
}

func TestPortReset(t *testing.T) {
	ports := setupPorts()

	t.Run("should empty the port set with zero ports", func(t *testing.T) {
		ports.Reset()

		if len(ports) > 0 {
			t.Error("set was not reset and still contains port records")
		}
	})
}
