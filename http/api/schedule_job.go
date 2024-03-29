package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	// "github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-offline"
)

// type ScheduleJobHandlerOptions defines a struct containing configuration options for the `ScheduleJobHandler` method.
type ScheduleJobHandlerOptions struct {
	// A `sfomuseum/go-offline.Database` instance to query for jobs.
	OfflineDatabase offline.Database
	// A `sfomuseum/go-offline.Queue` instance to schedule jobs.
	OfflineQueue offline.Queue
	// A `sfomuseum/go-http-auth.Authenticator` instance to use to restrict access.
	Authenticator auth.Authenticator
}

type ScheduleJobInput struct {
	// TBD...
}

// ScheduleJobHandler() returns an `http.Handler` instance that...
func ScheduleJobHandler(opts *ScheduleJobHandlerOptions) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		logger := slog.Default()

		acct, err := opts.Authenticator.GetAccountForRequest(req)

		if err != nil {
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

		enc_input, err := json.Marshal(input)

		if err != nil {
			logger.Error("Failed to encode input", "error", err)
			http.Error(rsp, "Failed to encode input", http.StatusBadRequest)
			return
		}

		job, err := offline.ScheduleJob(ctx, opts.OfflineDatabase, opts.OfflineQueue, acct.Name, string(enc_input))

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