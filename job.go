package offline

import (
	"context"
	"fmt"
	"time"
)

const (
	Pending Status = iota
	Queued
	Processing
	Completed
	Failed
)

type Status int

func (s Status) String() string {

	switch s {
	case Pending:
		return "pending"
	case Queued:
		return "queued"
	case Processing:
		return "processing"
	case Completed:
		return "completed"
	default:
		return "failed"
	}
}

type Job struct {
	Id           int64       `json:"id"`
	Status       Status      `json:"status"`
	Created      int64       `json:"created"`
	LastModified int64       `json:"lastmodified"`
	Instructions interface{} `json:"instruction"`
	Error        string      `json:"error,omitempty"`
}

func (job *Job) String() string {
	return fmt.Sprintf("%d (%v)", job.Id, job.Status)
}

func NewJob(ctx context.Context, instructions interface{}) (*Job, error) {

	id, err := NewJobId(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new job ID, %w", err)
	}

	now := time.Now()
	ts := now.Unix()

	job := &Job{
		Id:           id,
		Created:      ts,
		LastModified: ts,
		Status:       Pending,
		Instructions: instructions,
	}

	return job, nil
}
