package offline

import (
	"context"
	"fmt"
)

func ScheduleJob(ctx context.Context, offline_db Database, offline_q Queue, creator string, job_type string, instructions string) (*Job, error) {

	job, err := NewJob(ctx, creator, job_type, instructions)

	if err != nil {
		return nil, fmt.Errorf("Failed to create offline job, %w", err)
	}

	err = offline_db.AddJob(ctx, job)

	if err != nil {
		return nil, fmt.Errorf("Failed to add offline job, %w", err)
	}

	job.Status = Queued

	err = offline_db.UpdateJob(ctx, job)

	if err != nil {
		return nil, fmt.Errorf("Failed to update offline job status (to queued), %w", err)
	}

	err = offline_q.ScheduleJob(ctx, job.Id)

	if err != nil {

		job.Status = Pending

		status_err := offline_db.UpdateJob(ctx, job)

		if status_err != nil {
			return nil, fmt.Errorf("Failed to add offline job, %w. Also failed to update offline job status (to pending), %w", err, status_err)
		}

		return nil, fmt.Errorf("Failed to add offline job, %w", err)
	}

	return job, nil
}
