package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-offline/app/job/schedule"
)

func main() {

	ctx := context.Background()
	err := schedule.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to schedule job, %v", err)
	}
}
