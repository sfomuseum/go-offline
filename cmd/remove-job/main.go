package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-offline/app/job/remove"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := remove.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to add job, %v", err)
	}
}
