package main

import (
	"environment"
	"log/slog"
	"logging"
	"strconv"
)

func init() {
	logging.Initialize()

}

func main() {

	env := environment.Get()
	if env == nil {
		slog.Error("Failed to load the environment configuration.")
		return
	}

	portsAvailable, err := env.TestPortAvailability()
	if err != nil {
		slog.Error("Failed to test the port availability.")
		return
	}

	if !portsAvailable {
		slog.Error("Some ports are not available.")
		return
	}

	endpoint, error := environment.Get().FindEndpointByName("backbone")
	if error != nil {
		slog.Error("Failed to find the endpoint.")
		return
	}

	listener, error := endpoint.CreateListener()
	if error != nil {
		slog.Error("Failed to create the listener.")
		return
	}

	slog.Info("Listening on port: "+strconv.Itoa(int(endpoint.Port)), "endpoint", endpoint.Name)
	for {
		conn, error := listener.Accept()
		if error != nil {
			slog.Error("Failed to accept the connection.")
			return
		}

		// Http 1.1 response
		conn.Write([]byte("HTTP/1.1 200 OK\n"))
		conn.Write([]byte("Content-Type: text/plain\n"))
		conn.Write([]byte("\n"))
		conn.Write([]byte("Hello, World!"))
		conn.Close()

	}
}
