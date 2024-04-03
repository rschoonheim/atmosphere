package environment

import (
	"errors"
	"log/slog"
	"net"
	"strconv"
)

// FindEndpointByName - find the endpoint by name.
func (e *Environment) FindEndpointByName(name string) (*Endpoint, error) {
	for _, endpoint := range e.Endpoints {
		if endpoint.Name == name {
			return endpoint, nil
		}
	}
	return nil, errors.New("endpoint_not_found")
}

// TestPortAvailability - are all the ports available?
func (e *Environment) TestPortAvailability() (bool, error) {
	for _, endpoint := range e.Endpoints {
		if !testPortAvailability(endpoint.Port) {
			return false, errors.New("Port (" + string(endpoint.Port) + ") is not available.")
		}
	}

	return true, nil
}

// testPortAvailability - is the port available?
func testPortAvailability(port int32) bool {
	address := ":" + strconv.Itoa(int(port))
	listener, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("Failed to listen on address (" + address + ").")
		return false
	}
	listener.Close()
	return true
}
