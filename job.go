package offline

import (
	"context"
	// "encoding/json"
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
	Instructions interface{} `json:"instructions"`
	Results      string      `json:"results,omitempty"`
	Error        string      `json:"error,omitempty"`
}

type JobStatusResponse struct {
	JobId        int64  `json:"job_id"`
	Status       Status `json:"status"`
	LastModified int64  `json:"lastmodified"`
	Results      string `json:"results,omitempty"`
	Error        string `json:"error,omitempty"`
}

func NewJob(ctx context.Context, instructions interface{}) (*Job, error) {

	/*
		enc_instructions, err := json.Marshal(instructions)

		if err != nil {
			return nil, fmt.Errorf("Failed to marshal instructions, %w", err)
		}
	*/

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

	status_rsp := &JobStatusResponse{
		JobId:        job.Id,
		Status:       job.Status,
		LastModified: job.LastModified,
	}

	return status_rsp
}
