package sl

import (
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func Obj(value any) slog.Attr {
	return slog.Attr{Key: "object", Value: slog.AnyValue(value)}
}
