package main

import (
	"context"
	"flag"
	"github.com/sfomuseum/go-offline"
	"log"
)

func main() {

	database_uri := flag.String("database-uri", "awsdynamodb://OfflineJobs?partition_key=Id&local=true", "")

	job_id := flag.Int64("job-id", 0, "")

	flag.Parse()

	ctx := context.Background()

	db, err := offline.NewDatabase(ctx, *database_uri)

	if err != nil {
		log.Fatalf("Failed to create offline database, %v", err)
	}

	job, err := db.GetJob(ctx, *job_id)

	if err != nil {
		log.Fatalf("Failed to get job, %v", err)
	}

	err = db.RemoveJob(ctx, job)

	if err != nil {
		log.Fatalf("Failed to remove job, %v", err)
	}
}
