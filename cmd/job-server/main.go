package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-offline/app/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run job server, %v", err)
	}
}
