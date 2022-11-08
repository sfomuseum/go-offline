package api

import (
	"encoding/json"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-offline"
	"net/http"
)

type JobStatusHandlerOptions struct {
	Database      offline.Database
	Authenticator auth.Authenticator
}

type JobStatusResponse struct {
	JobId        int64          `json:"job_id"`
	Status       offline.Status `json:"status"`
	LastModified int64          `json:"lastmodified"`
}

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

		status_rsp := JobStatusResponse{
			JobId:        job.Id,
			Status:       job.Status,
			LastModified: job.LastModified,
		}

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		err = enc.Encode(status_rsp)

		if err != nil {
			http.Error(rsp, "Server error", http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn)
}
