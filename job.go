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

type Job struct {
	Id           int64       `json:"id"`
	Status       Status      `json:"status"`
	Created      int64       `json:"created"`
	LastModified int64       `json:"lastmodified"`
	Instructions interface{} `json:"instruction"`
	Error        string      `json:"error,omitempty"`
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
