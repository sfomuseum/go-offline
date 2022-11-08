package offline

import (
	"context"
	"testing"
	_ "gocloud.dev/docstore/memdocstore"
)

func TestDocstoreDatabase(t *testing.T) {

	ctx := context.Background()
	
	db_uri := "mem://offline/Id"

	db, err := NewDatabase(ctx, db_uri)

	if err != nil {
		t.Fatalf("Failed to create new database, %v", err)
	}

	instructions := map[string]interface{}{
		"name": "testing",
		"id":   1234,
	}
	
	job, err := NewJob(ctx, instructions)

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
