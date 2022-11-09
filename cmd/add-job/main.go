package main

import (
	_ "gocloud.dev/docstore/awsdynamodb"
)

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-offline"
	"log"
)

func main() {

	database_uri := flag.String("database-uri", "awsdynamodb://OfflineJobs?partition_key=Id&local=true", "")
	instructions := flag.String("instructions", "", "")

	flag.Parse()

	ctx := context.Background()

	db, err := offline.NewDatabase(ctx, *database_uri)

	if err != nil {
		log.Fatalf("Failed to create offline database, %v", err)
	}

	var data interface{}

	err = json.Unmarshal([]byte(*instructions), &data)

	if err != nil {
		log.Fatalf("Failed to unmarshal instructions, %v", err)
	}

	job, err := offline.NewJob(ctx, data)

	if err != nil {
		log.Fatalf("Failed to create new job, %v", err)
	}

	err = db.AddJob(ctx, job)

	if err != nil {
		log.Fatalf("Failed to add job, %v", err)
	}

	fmt.Println(job.Id)
}
