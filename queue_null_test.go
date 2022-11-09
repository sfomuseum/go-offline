package offline

import (
	"context"
	"testing"
)

func TestQueueJob(t *testing.T) {

	ctx := context.Background()

	q_uri := "null://"

	q, err := NewQueue(ctx, q_uri)

	if err != nil {
		t.Fatalf("Failed to create new queue, %v", err)
	}

	job_id, err := NewJobId(ctx)

	if err != nil {
		t.Fatalf("Failed to create new job ID, %v", err)
	}

	err = q.ScheduleJob(ctx, job_id)

	if err != nil {
		t.Fatalf("Failed to schedule job, %v", err)
	}

	err = q.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close queue, %v", err)
	}
}
