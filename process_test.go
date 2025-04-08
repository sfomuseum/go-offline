package offline

import (
	"context"
	"encoding/json"
	"testing"
)

func TestProcessJob(t *testing.T) {

	ctx := context.Background()

	db_uri := "syncmap://"

	db, err := NewDatabase(ctx, db_uri)

	if err != nil {
		t.Fatalf("Failed to create new database, %v", err)
	}

	job_type := "testing"

	instructions := map[string]interface{}{
		"name": "testing",
		"id":   1234,
	}

	enc_instructions, err := json.Marshal(instructions)

	if err != nil {
		t.Fatalf("Failed to marshal instructions, %v", err)
	}

	str_instructions := string(enc_instructions)

	job, err := NewJob(ctx, "testing", job_type, str_instructions)

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

	process_cb := func(ctx context.Context, job *Job) (string, error) {
		return "OK", nil
	}

	process_opts := &ProcessJobOptions{
		Database: db,
		Callback: process_cb,
		JobId:    job.Id,
	}

	err = ProcessJob(ctx, process_opts)

	if err != nil {
		t.Fatalf("Failed to process job, %v", err)
	}

	job2, err := db.GetJob(ctx, job.Id)

	if err != nil {
		t.Fatalf("Failed to retrieve job %d, %v", job.Id, err)
	}

	if job2.Results != "OK" {
		t.Fatalf("Unexpected job results, %s", job2.Results)
	}
}
