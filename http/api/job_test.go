package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-offline"
	"net/http"
	"testing"
)

func TestJobStatusHandler(t *testing.T) {

	ctx := context.Background()

	db_uri := "syncmap://"

	db, err := offline.NewDatabase(ctx, db_uri)

	if err != nil {
		t.Fatalf("Failed to create new database, %v", err)
	}

	authenticator, err := auth.NewAuthenticator(ctx, "null://")

	if err != nil {
		t.Fatalf("Failed to create new authenticator, %v", err)
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

	job, err := offline.NewJob(ctx, str_instructions)

	if err != nil {
		t.Fatalf("Failed to create new job, %v", err)
	}

	err = db.AddJob(ctx, job)

	if err != nil {
		t.Fatalf("Failed to add job, %v", err)
	}

	handler_opts := &JobStatusHandlerOptions{
		Database:      db,
		Authenticator: authenticator,
	}

	handler := JobStatusHandler(handler_opts)

	go func() {

		err := http.ListenAndServe("localhost:8080", handler)

		if err != nil {
			t.Fatalf("Failed to serve requests, %v", err)
		}
	}()

	url := fmt.Sprintf("http://localhost:8080?id=%d", job.Id)

	rsp, err := http.Get(url)

	if err != nil {
		t.Fatalf("Failed to get '%s', %v", url, err)
	}

	var status_rsp offline.JobStatusResponse

	dec := json.NewDecoder(rsp.Body)
	err = dec.Decode(&status_rsp)

	if err != nil {
		t.Fatalf("Failed to decode response, %v", err)
	}

	if status_rsp.JobId != job.Id {
		t.Fatalf("Unexpected job ID, %d", status_rsp.JobId)
	}
}
