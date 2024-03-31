package main

import (
	"environment"
	"log/slog"
	"logging"
	"os"
)

func init() {
	logging.Initialize()

	// Make sure that each endpoint uses a unique port.
	endpointPortConflict := environment.Get().HasDuplicatePorts()
	if endpointPortConflict > 0 {
		slog.Error("Could not start the application, there are multiple endpoints with the same port.", "port", endpointPortConflict)
		os.Exit(1)
	}
}

func main() {

	for _, arg := range environment.Get().Endpoints {
		slog.Info("Endpoint", arg.Name, arg.Port, arg.TlsCert.Cert, arg.TlsCert.Key)

		_, err := arg.TlsCert.Exists()
		if err != nil {
			slog.Error("Failed to validate TLS certificate and key files.", "error", err)
			os.Exit(1)
		}
	}
}
