package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-offline/app/job/get"
)

func main() {

	ctx := context.Background()
	err := get.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to add job, %v", err)
	}
}
