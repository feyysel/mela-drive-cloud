package slogx

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func Init(debug bool) {
	var handler slog.Handler

	if debug {
		handler = tint.NewTextHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: "[15:04:05]",
		})
	} else {
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	slog.SetDefault(slog.New(handler))
}
