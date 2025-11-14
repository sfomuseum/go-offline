package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-http/v4/auth"
	"github.com/aaronland/go-http/v4/sanitize"
	"github.com/sfomuseum/go-offline"
	off_http "github.com/sfomuseum/go-offline/http"
)

// type JobStatusHandlerOptions defines a struct containing configuration options for the `JobStatusHandler` method.
type JobStatusHandlerOptions struct {
	// A `sfomuseum/go-offline.Database` instance to query for jobs.
	OfflineDatabase offline.Database
	// A `sfomuseum/go-http-auth.Authenticator` instance to use to restrict access.
	Authenticator auth.Authenticator
}

// JobStatusHandler() returns an `http.Handler` instance that will return a JSON-encoded `JobStatusResponse`
// for a job identified by a HTTP GET "id" query parameter.
func JobStatusHandler(opts *JobStatusHandlerOptions) http.Handler {

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

		id, err := sanitize.GetInt64(req, "id")

		if err != nil {
			http.Error(rsp, "Invalid id", http.StatusBadRequest)
			return
		}

		logger = logger.With("job id", id)

		job, err := opts.OfflineDatabase.GetJob(ctx, id)

		if err != nil {
			logger.Error("Failed to retrieve job", "error", err)
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(job.AsStatusResponse())

		if err != nil {
			logger.Error("Failed to encode job response", "error", err)
			http.Error(rsp, "Server error", http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn)
}
