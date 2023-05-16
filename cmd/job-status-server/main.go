package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-offline/app/status/server"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := server.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to add job, %v", err)
	}
}
