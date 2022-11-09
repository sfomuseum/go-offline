package main

import (
	"context"
	"github.com/sfomuseum/go-offline/app/job/get"
	"log"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := get.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to add job, %v", err)
	}
}
