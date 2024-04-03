package environment

import (
	"net"
	"strconv"
)

// CreateListener - create a listener based on the endpoint configuration.
func (e *Endpoint) CreateListener() (net.Listener, error) {
	address := ":" + strconv.Itoa(int(e.Port))
	return net.Listen("tcp", address)
}
