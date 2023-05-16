package api

import (
	"encoding/json"
	"net/http"

	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-offline"
)

// type JobStatusHandlerOptions defines a struct containing configuration options for the `JobStatusHandler` method.
type JobStatusHandlerOptions struct {
	// A `sfomuseum/go-offline.Database` instance to query for jobs.
	Database offline.Database
	// A `sfomuseum/go-http-auth.Authenticator` instance to use to restrict access.
	Authenticator auth.Authenticator
}

// JobStatusHandler() returns an `http.Handler` instance that will return a JSON-encoded `JobStatusResponse`
// for a job identified by a HTTP GET "id" query parameter.
func JobStatusHandler(opts *JobStatusHandlerOptions) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		_, err := opts.Authenticator.GetAccountForRequest(req)

		if err != nil {
			http.Error(rsp, "Not authorized", http.StatusUnauthorized)
			return
		}

		id, err := sanitize.GetInt64(req, "id")

		if err != nil {
			http.Error(rsp, "Invalid id", http.StatusBadRequest)
			return
		}

		job, err := opts.Database.GetJob(ctx, id)

		if err != nil {
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(job.AsStatusResponse())

		if err != nil {
			http.Error(rsp, "Server error", http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn)
}
