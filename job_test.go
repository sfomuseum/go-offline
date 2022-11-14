package offline

import (
	"context"
	"encoding/json"
	"github.com/tidwall/gjson"
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

func TestJobStatusResponse(t *testing.T) {

	ctx := context.Background()

	j, err := NewJob(ctx, "testing")

	if err != nil {
		t.Fatalf("Failed to create new job, %v", err)
	}

	j.Id = 1592265428804046848

	s := j.AsStatusResponse()

	if s.JobId != j.Id {
		t.Fatalf("Invalid job status ID, %d", s.JobId)
	}

	enc_s, err := json.Marshal(s)

	if err != nil {
		t.Fatalf("Failed to marshal job status, %v", err)
	}

	id_rsp := gjson.GetBytes(enc_s, "job_id")

	if id_rsp.Int() != j.Id {
		t.Fatalf("Invalid job status ID (encoded), %d", id_rsp.Int())
	}
}
