package offline

import (
	"context"
	"fmt"
	"log"
)

type ProcessJobCallbackFunc func(context.Context, *Job) (string, error)

type ProcessJobOptions struct {
	Database Database
	Logger   *log.Logger
	Callback ProcessJobCallbackFunc
	JobId    int64
}

func ProcessJob(ctx context.Context, opts *ProcessJobOptions) error {

	offline_db := opts.Database
	job_id := opts.JobId

	logger := opts.Logger

	job, err := offline_db.GetJob(ctx, job_id)

	if err != nil {
		return fmt.Errorf("Failed to retrieve job %d, %w", job_id, err)
	}

	if job.Status != Queued {
		logger.Printf("Job %d not marked as queued (%s), skipping", job_id, job.Status)
		return nil
	}

	job.Status = Processing

	err = offline_db.UpdateJob(ctx, job)

	if err != nil {
		return fmt.Errorf("Failed to update status for job %d, %w", job_id, err)
	}

	var final_status Status
	var final_results string
	var final_error error

	defer func() {

		job.Status = final_status

		if final_error != nil {
			job.Error = final_error.Error()
		} else {
			job.Results = final_results
		}

		err = offline_db.UpdateJob(ctx, job)

		if err != nil {
			logger.Printf("Failed to update final status (%s) for job %d, %v", final_status, job_id, err)
		}
	}()

	results, err := opts.Callback(ctx, job)

	if err != nil {

		final_status = Failed
		final_error = err

		return fmt.Errorf("Callback function failed, %w", err)
	}

	final_status = Completed
	final_results = results

	return nil
}
