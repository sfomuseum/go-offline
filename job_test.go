package offline

import (
	"context"
	"encoding/json"
	"testing"
)

func TestNewJob(t *testing.T) {

	ctx := context.Background()

	instructions := map[string]interface{}{
		"name": "testing",
		"id":   1234,
	}

	enc_instructions, err := json.Marshal(instructions)

	if err != nil {
		t.Fatalf("Failed to marshal instructions, %v", err)
	}

	str_instructions := string(enc_instructions)

	j, err := NewJob(ctx, str_instructions)

	if err != nil {
		t.Fatalf("Failed to create new job, %v", err)
	}

	if j.Status != Pending {
		t.Fatalf("Unexpected status: %v", j.Status)
	}
}
