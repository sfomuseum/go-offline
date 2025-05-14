package get

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-offline"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	offline_db, err := offline.NewDatabase(ctx, database_uri)

	if err != nil {
		return fmt.Errorf("Failed to create offline database, %w", err)
	}

	offline_q, err := offline.NewQueue(ctx, queue_uri)

	if err != nil {
		return fmt.Errorf("Failed to create offline queue, %w", err)
	}

	job, err := offline_db.GetJob(ctx, job_id)

	if err != nil {
		return fmt.Errorf("Failed to get job, %w", err)
	}

	if job.Status != offline.Pending {
		return fmt.Errorf("Job status is not pending (%d)", job.Status)
	}

	job.Status = offline.Queued

	err = offline_db.UpdateJob(ctx, job)

	if err != nil {
		return fmt.Errorf("Failed to update offline job status (to queued), %w", err)
	}

	err = offline_q.ScheduleJob(ctx, job.Id)

	if err != nil {

		job.Status = offline.Pending

		status_err := offline_db.UpdateJob(ctx, job)

		if status_err != nil {
			return fmt.Errorf("Failed to add offline job, %w. Also failed to update offline job status (to pending), %w", err, status_err)
		}

		return fmt.Errorf("Failed to add offline job, %w", err)
	}

	// Wait for job to complete?

	return nil
}
