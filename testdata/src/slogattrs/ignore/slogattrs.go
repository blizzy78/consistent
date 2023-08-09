package slogattrs

import "log/slog"

func defaultLogger() {
	slog.Info("test", "a", "b", "c", "d")
	slog.Info("test", "a", "b", slog.String("c", "d"))
	slog.Info("test", slog.String("a", "b"), slog.String("c", "d"))
}

func logger(logger *slog.Logger) {
	logger.Info("test", "a", "b", "c", "d")
	logger.Info("test", "a", "b", slog.String("c", "d"))
	logger.Info("test", slog.String("a", "b"), slog.String("c", "d"))
}

func other() {
	println("test", "a", "b", "c", "d")
	println("test", "a", "b", slog.String("c", "d"))
	println("test", slog.String("a", "b"), slog.String("c", "d"))
}
