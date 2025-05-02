package utils

import (
	"io"
	"log/slog"
	"path/filepath"
)

func Mylog(w io.Writer, level slog.Level) *slog.Logger {
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
