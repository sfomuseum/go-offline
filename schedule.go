package offline

import (
	"context"
	"fmt"
)

func ScheduleJob(ctx context.Context, offline_db Database, offline_q Queue, creator string, instructions string) (*Job, error) {

	job, err := NewJob(ctx, creator, instructions)

	if err != nil {
		return nil, fmt.Errorf("Failed to create offline job, %v", err)
	}

	err = offline_db.AddJob(ctx, job)

	if err != nil {
		return nil, fmt.Errorf("Failed to add offline job, %v", err)
	}

	err = offline_q.ScheduleJob(ctx, job.Id)

	if err != nil {
		return nil, fmt.Errorf("Failed to add offline job, %v", err)
	}

	job.Status = Queued

	err = offline_db.UpdateJob(ctx, job)

	if err != nil {
		return nil, fmt.Errorf("Failed to update offline job status, %v", err)
	}

	return job, nil
}
