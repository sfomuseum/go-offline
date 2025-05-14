package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-offline/app/job/add"
)

func main() {

	ctx := context.Background()
	err := add.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to add job, %v", err)
	}
}
