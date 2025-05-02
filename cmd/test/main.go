package main

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func makeLog(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource: true,
		Level:     &level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}))
}

func main() {
	logger := makeLog(os.Stderr, slog.LevelDebug)
	logger.Info("Hello World!")
}
