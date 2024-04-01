package main

import (
	"log/slog"
	"os"

	"github.com/ostapneko/muso-matrix/internal/image"
	"github.com/ostapneko/muso-matrix/internal/muso"
)

func main() {
	slog.Info("Starting application...")
	result, err := muso.GetTrackInfo()

	if err != nil {
		slog.Error("Failed to get track info", err)
		os.Exit(1)
	}

	slog.Info("Track info", "result", result)

	image.Draw(result, 10)
}
