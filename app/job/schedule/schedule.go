package schedule

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-offline"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	logger := slog.Default()
	logger = logger.With("job id", job_id)

	offline_db, err := offline.NewDatabase(ctx, offline_database_uri)

	if err != nil {
		return fmt.Errorf("Failed to create offline database, %w", err)
	}

	// START OF put me in a function with well-defined types etc.

	q_mux := make(map[string]offline.Queue)

	for _, kv := range offline_queue_uris {

		job_type := kv.Key()
		offline_uri := kv.Value().(string)

		_, exists := q_mux[job_type]

		if exists {
			return fmt.Errorf("Multiple values for '%s' job type", job_type)
		}

		offline_uri, err := url.QueryUnescape(offline_uri)

		if err != nil {
			return fmt.Errorf("Failed to unescape URI '%s' for job '%s', %w", offline_uri, job_type, err)
		}

		offline_q, err := offline.NewQueue(ctx, offline_uri)

		if err != nil {
			return fmt.Errorf("Failed to instantiate offline queue for '%s', %w", job_type, err)
		}

		q_mux[job_type] = offline_q
	}

	// END OF put me in a function with well-defined types etc.

	job, err := offline_db.GetJob(ctx, job_id)

	if err != nil {
		return fmt.Errorf("Failed to get job, %w", err)
	}

	offline_q, ok := q_mux[job.Type]

	if !ok {
		return fmt.Errorf("Failed to derive offline queue for job type (%s)", job.Type)
	}

	if job.Status != offline.Pending {
		return fmt.Errorf("Job status is not pending (%d)", job.Status)
	}

	logger.Debug("Set job status to queued")

	job.Status = offline.Queued

	err = offline_db.UpdateJob(ctx, job)

	if err != nil {
		return fmt.Errorf("Failed to update offline job status (to queued), %w", err)
	}

	logger.Debug("Schedule job")

	err = offline_q.ScheduleJob(ctx, job.Id)

	if err != nil {

		logger.Debug("Set job status to pending")

		job.Status = offline.Pending

		status_err := offline_db.UpdateJob(ctx, job)

		if status_err != nil {
			return fmt.Errorf("Failed to schedule offline job, %w. Also failed to update offline job status (to pending), %w", err, status_err)
		}

		return fmt.Errorf("Failed to schedule offline job, %w", err)
	}

	// Wait for job to complete?

	return nil
}
