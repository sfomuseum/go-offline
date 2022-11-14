package offline

import (
	"context"
	"fmt"
	"strconv"
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
	Id           int64  `json:"id"`
	Status       Status `json:"status"`
	Created      int64  `json:"created"`
	LastModified int64  `json:"lastmodified"`
	Instructions string `json:"instructions"`
	Results      string `json:"results,omitempty"`
	Error        string `json:"error,omitempty"`
}

type JobStatusResponse struct {
	JobId        string `json:"job_id"` // because JavaScript 58-bit integer hoohah
	Status       string `json:"status"`
	LastModified int64  `json:"lastmodified"`
	Results      string `json:"results,omitempty"`
	Error        string `json:"error,omitempty"`
}

func NewJob(ctx context.Context, instructions string) (*Job, error) {

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

func (job *Job) String() string {
	return fmt.Sprintf("%d (%v)", job.Id, job.Status)
}

func (job *Job) AsStatusResponse() *JobStatusResponse {

	str_id := strconv.FormatInt(job.Id, 10)

	status_rsp := &JobStatusResponse{
		JobId:        str_id,
		Status:       job.Status.String(),
		LastModified: job.LastModified,
	}

	return status_rsp
}
