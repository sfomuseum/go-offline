package offline

import (
	"context"
	"fmt"
	"log/slog"
)

type ProcessJobCallbackFunc func(context.Context, *Job) (string, error)

type ProcessJobOptions struct {
	Database Database
	Callback ProcessJobCallbackFunc
	JobId    int64
}

func ProcessJob(ctx context.Context, opts *ProcessJobOptions) error {

	offline_db := opts.Database
	job_id := opts.JobId

	logger := slog.Default()
	logger = logger.With("job id", job_id)

	logger.Debug("Fetch job")
	
	job, err := offline_db.GetJob(ctx, job_id)

	if err != nil {
		logger.Error("Failed to retrieve job", "error", err)
		return fmt.Errorf("Failed to retrieve job %d, %w", job_id, err)
	}

	if job.Status != Queued {
		logger.Warn("Job not marked as queued", "status", job.Status)
		return nil
	}

	logger.Debug("Mark job as processing")
	job.Status = Processing

	err = offline_db.UpdateJob(ctx, job)

	if err != nil {
		logger.Error("Failed to update job status", "error", err)
		return fmt.Errorf("Failed to update status for job %d, %w", job_id, err)
	}

	var final_status Status
	var final_results string
	var final_error error

	defer func() {

		logger.Debug("Update job with final status", "status", final_status)
		job.Status = final_status

		if final_error != nil {
			job.Error = final_error.Error()
		} else {
			job.Results = final_results
		}

		err = offline_db.UpdateJob(ctx, job)

		if err != nil {
			logger.Error("Failed to update final status", "status", final_status, "error", err)
		}
	}()

	logger.Debug("Dispatch job callback")
	results, err := opts.Callback(ctx, job)

	if err != nil {

		final_status = Failed
		final_error = err

		logger.Error("Final callback failed", "error", err)
		return fmt.Errorf("Callback function failed, %w", err)
	}

	final_status = Completed
	final_results = results

	return nil
}
