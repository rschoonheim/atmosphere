package logging

import (
	"log/slog"
	"os"
)

// Initialize - configures the standard slog logger.
func Initialize() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
