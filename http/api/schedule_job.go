package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-offline"
	off_http "github.com/sfomuseum/go-offline/http"
)

// type ScheduleJobHandlerOptions defines a struct containing configuration options for the `ScheduleJobHandler` method.
type ScheduleJobHandlerOptions struct {
	// A `sfomuseum/go-offline.Database` instance to query for jobs.
	OfflineDatabase offline.Database
	// A `sfomuseum/go-offline.Queue` instance to schedule jobs.
	// OfflineQueue offline.Queue
	// A `sfomuseum/go-http-auth.Authenticator` instance to use to restrict access.
	OfflineQueueMux map[string]offline.Queue
	Authenticator   auth.Authenticator
}

type ScheduleJobInput struct {
	Type         string      `json:"type"`
	Instructions interface{} `json:"instructions"`
}

// ScheduleJobHandler() returns an `http.Handler` instance that...
func ScheduleJobHandler(opts *ScheduleJobHandlerOptions) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := slog.Default()

		logger = off_http.LoggerWithRequest(req, logger)

		acct, err := opts.Authenticator.GetAccountForRequest(req)

		if err != nil {
			logger.Error("Not authorized", "error", err)
			http.Error(rsp, "Not authorized", http.StatusUnauthorized)
			return
		}

		logger = logger.With("account", acct.Name)

		var input *ScheduleJobInput

		dec := json.NewDecoder(req.Body)
		err = dec.Decode(&input)

		if err != nil {
			logger.Error("Failed to decode request body", "error", err)
			http.Error(rsp, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		job_type := input.Type

		if job_type == "" {
			logger.Error("Missing job type")
			http.Error(rsp, "Bad request", http.StatusBadRequest)
			return
		}

		offline_q, exists := opts.OfflineQueueMux[job_type]

		if !exists {
			offline_q, exists = opts.OfflineQueueMux["*"]
		}

		if !exists {
			logger.Error("Failed to derive queue for job", "type", job_type)
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		logger = logger.With("job type", job_type)

		enc_instructions, err := json.Marshal(input.Instructions)

		if err != nil {
			logger.Error("Failed to encode instructions", "error", err)
			http.Error(rsp, "Failed to encode instructions", http.StatusBadRequest)
			return
		}

		job, err := offline.ScheduleJob(ctx, opts.OfflineDatabase, offline_q, acct.Name(), job_type, string(enc_instructions))

		if err != nil {
			logger.Error("Failed to schedule update for offline job", "error", err)
			http.Error(rsp, "Failed to schedule update for offline job", http.StatusInternalServerError)
			return
		}

		logger = logger.With("job id", job.Id)

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(job.AsStatusResponse())

		if err != nil {
			logger.Error("Failed to encode job status response", "error", err)
			http.Error(rsp, "Failed to encode job status response", http.StatusInternalServerError)
			return
		}

		logger.Info("Job successfully scheduled")
	}

	return http.HandlerFunc(fn)
}
