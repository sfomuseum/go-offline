package offline

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	_ "gocloud.dev/docstore/memdocstore"

	"encoding/json"
)

func TestDocstoreDatabase(t *testing.T) {

	ctx := context.Background()

	db_uri := "mem://offline/Id"

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

	job, err := NewJob(ctx, "testing", job_type, string(enc_instructions))

	if err != nil {
		t.Fatalf("Failed to create new job, %v", err)
	}

	err = db.AddJob(ctx, job)

	if err != nil {
		t.Fatalf("Failed to add job, %v", err)
	}

	job, err = db.GetJob(ctx, job.Id)

	if err != nil {
		t.Fatalf("Failed to retrieve job, %v", err)
	}

	job.Status = Processing

	err = db.UpdateJob(ctx, job)

	if err != nil {
		t.Fatalf("Failed to update job, %v", err)
	}

	job, err = db.GetJob(ctx, job.Id)

	if err != nil {
		t.Fatalf("Failed to retrieve job (again), %v", err)
	}

	if job.Status != Processing {
		t.Fatalf("Expected job to be processed but is: %v", job.Status)
	}

	err = db.RemoveJob(ctx, job)

	if err != nil {
		t.Fatalf("Failed to delete job")
	}

	job, _ = db.GetJob(ctx, job.Id)

	if job != nil {
		t.Fatalf("Expected to not find job (after deleting) but it's still there")
	}
}

func TestGocloudPruneAndListJobs(t *testing.T) {

	ctx := context.Background()

	db_uri := "mem://offline/Id"

	db, err := NewDatabase(ctx, db_uri)

	if err != nil {
		t.Fatalf("Failed to create new database, %v", err)
	}

	job_type := "testing"

	instructions := map[string]interface{}{
		"name": "testing",
		"id":   1234,
	}

	for i := 0; i < 5; i++ {

		enc_instructions, err := json.Marshal(instructions)

		if err != nil {
			t.Fatalf("Failed to marshal instructions, %v", err)
		}

		job, err := NewJob(ctx, "testing", job_type, string(enc_instructions))

		if err != nil {
			t.Fatalf("Failed to create new job, %v", err)
		}

		err = db.AddJob(ctx, job)

		if err != nil {
			t.Fatalf("Failed to add job, %v", err)
		}
	}

	now := time.Now()
	ts := now.Unix()

	err = db.PruneJobs(ctx, Pending, ts)

	if err != nil {
		t.Fatalf("Failed to prune jobs, %v", err)
	}

	count := int32(0)

	list_cb := func(ctx context.Context, job *Job) error {

		atomic.AddInt32(&count, 1)
		return nil
	}

	err = db.ListJobs(ctx, list_cb)

	if err != nil {
		t.Fatalf("Failed to list jobs, %v", err)
	}

	if count != 0 {
		t.Fatalf("Expected job count to be 0, not %d", count)
	}

	err = db.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close collection, %v", err)
	}

}
