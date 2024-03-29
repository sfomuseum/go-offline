package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/sfomuseum/go-offline/app/server"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		logger.Error("Failed to add job", "error", err)
		os.Exit(1)
	}
}
