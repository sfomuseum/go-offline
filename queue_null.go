package offline

import (
	"context"
)

type NullQueue struct {
	Queue
}

func init() {
	ctx := context.Background()
	RegisterQueue(ctx, "null", NewNullQueue)
}

func NewNullQueue(ctx context.Context, uri string) (Queue, error) {
	q := &NullQueue{}
	return q, nil
}

func (q *NullQueue) ScheduleJob(ctx context.Context, job_id int64) error {
	return nil
}

func (q *NullQueue) Close(ctx context.Context) error {
	return nil
}
