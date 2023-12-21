package main

import (
	"os"
	"log/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello")
}