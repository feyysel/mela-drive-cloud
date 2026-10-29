package main

import (
	"log/slog"
	"os"

	"github.com/feyysel/mela-drive-cloud/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error("server failed", "err", err)
		os.Exit(1)
	}
}
