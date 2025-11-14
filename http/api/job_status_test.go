package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/aaronland/go-http/v4/auth"
	"github.com/sfomuseum/go-offline"
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

	job, err := offline.NewJob(ctx, "testing", job_type, str_instructions)

	if err != nil {
		t.Fatalf("Failed to create new job, %v", err)
	}

	err = db.AddJob(ctx, job)

	if err != nil {
		t.Fatalf("Failed to add job, %v", err)
	}

	handler_opts := &JobStatusHandlerOptions{
		OfflineDatabase: db,
		Authenticator:   authenticator,
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

	if status_rsp.JobId != strconv.FormatInt(job.Id, 10) {
		t.Fatalf("Unexpected job ID, %s", status_rsp.JobId)
	}
}
