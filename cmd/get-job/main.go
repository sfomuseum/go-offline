package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/sfomuseum/go-offline"
	"log"
	"os"
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

	log.Println("GET", *job_id)

	job, err := db.GetJob(ctx, *job_id)

	if err != nil {
		log.Fatalf("Failed to add job, %v", err)
	}

	dec := json.NewDecoder(os.Stdout)
	err = dec.Decode(job)

	if err != nil {
		log.Fatalf("Failed to decode job, %v", err)
	}

}
