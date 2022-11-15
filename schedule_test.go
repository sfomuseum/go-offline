package offline

import (
	"context"
	"encoding/json"
	"testing"
)

func TestScheduleJob(t *testing.T) {

	ctx := context.Background()

	db_uri := "syncmap://"

	db, err := NewDatabase(ctx, db_uri)

	if err != nil {
		t.Fatalf("Failed to create new database, %v", err)
	}

	q_uri := "null://"

	q, err := NewQueue(ctx, q_uri)

	if err != nil {
		t.Fatalf("Failed to create new queue, %v", err)
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

	_, err = ScheduleJob(ctx, db, q, "testing", str_instructions)

	if err != nil {
		t.Fatalf("Failed to schedule job, %v", err)
	}
}
