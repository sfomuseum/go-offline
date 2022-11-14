package offline

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestProcessJob(t *testing.T) {

	ctx := context.Background()

	db_uri := "syncmap://"

	db, err := NewDatabase(ctx, db_uri)

	if err != nil {
		t.Fatalf("Failed to create new database, %v", err)
	}

	instructions := map[string]interface{}{
		"name": "testing",
		"id":   1234,
	}

	enc_instructions, err := json.Marshal(instructions)

	if err != nil {
		t.Fatalf("Failed to marshal instructions, %v", err)
	}

	str_instructions := string(enc_instructions)

	job, err := NewJob(ctx, str_instructions)

	if err != nil {
		t.Fatalf("Failed to create new job, %v", err)
	}

	err = db.AddJob(ctx, job)

	if err != nil {
		t.Fatalf("Failed to add job, %v", err)
	}

	job.Status = Queued

	err = db.UpdateJob(ctx, job)

	if err != nil {
		t.Fatalf("Failed to update job, %v", err)
	}

	process_cb := func(ctx context.Context, job *Job) error {
		return nil
	}

	process_logger := log.Default()

	process_opts := &ProcessJobOptions{
		Database: db,
		Logger:   process_logger,
		Callback: process_cb,
		JobId:    job.Id,
	}

	err = ProcessJob(ctx, process_opts)

	if err != nil {
		t.Fatalf("Failed to process job, %v", err)
	}
}
