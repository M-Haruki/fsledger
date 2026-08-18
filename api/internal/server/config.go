package server

import "log/slog"

type Config struct {
	IsDev    bool
	LogLevel slog.Leveler
}
