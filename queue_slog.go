package offline

import (
	"context"
	"log/slog"
)

type SlogQueue struct {
	Queue
}

func init() {
	ctx := context.Background()
	RegisterQueue(ctx, "slog", NewSlogQueue)
}

func NewSlogQueue(ctx context.Context, uri string) (Queue, error) {
	q := &SlogQueue{}
	return q, nil
}

func (q *SlogQueue) ScheduleJob(ctx context.Context, job_id int64) error {
	slog.Info("Schedule job", "id", job_id)
	return nil
}

func (q *SlogQueue) Close(ctx context.Context) error {
	return nil
}
