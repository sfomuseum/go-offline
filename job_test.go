package offline

import (
	"context"
	"testing"
)

func TestNewJob(t *testing.T) {

	ctx := context.Background()

	instructions := map[string]interface{}{
		"name": "testing",
		"id":   1234,
	}

	j, err := NewJob(ctx, instructions)

	if err != nil {
		t.Fatalf("Failed to create new job, %v", err)
	}

	if j.Status != Pending {
		t.Fatalf("Unexpected status: %v", j.Status)
	}
}
