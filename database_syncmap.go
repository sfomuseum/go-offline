package offline

import (
	"context"
	"fmt"
	"sync"
)

// SyncMapDatabase implements the `Database` interface storing and retrieving jobs from an internal `sync.Map` instance.
type SyncMapDatabase struct {
	Database
	sync_map *sync.Map
}

func init() {
	ctx := context.Background()
	RegisterDatabase(ctx, "syncmap", NewSyncMapDatabase)
}

// NewSyncMapDatabase returns a new `SyncMapDatabase` instance configured by 'uri' which is
// expected to take the form of:
//
//	syncmap://
func NewSyncMapDatabase(ctx context.Context, uri string) (Database, error) {

	sync_map := new(sync.Map)

	db := &SyncMapDatabase{
		sync_map: sync_map,
	}

	return db, nil
}

// AddJob() stores 'job' in 'db'.
func (db *SyncMapDatabase) AddJob(ctx context.Context, job *Job) error {
	db.sync_map.Store(job.Id, job)
	return nil
}

// GetJob() returns a `Job` instance identified by 'job_id' from 'db'.
func (db *SyncMapDatabase) GetJob(ctx context.Context, job_id int64) (*Job, error) {
	v, ok := db.sync_map.Load(job_id)

	if !ok {
		return nil, fmt.Errorf("Not found")
	}

	return v.(*Job), nil
}

// UpdateJob() stores the updated version of 'job' in 'db'.
func (db *SyncMapDatabase) UpdateJob(ctx context.Context, job *Job) error {
	db.sync_map.Store(job.Id, job)
	return nil
}

// RemoveJob() removes 'job' from 'db'.
func (db *SyncMapDatabase) RemoveJob(ctx context.Context, job *Job) error {
	db.sync_map.Delete(job.Id)
	return nil
}

// PruneJobs() removed jobs in 'db' whose status matches 'status' and whose last modified time is less that 'lastmodified'.
func (db *SyncMapDatabase) PruneJobs(ctx context.Context, status Status, lastmodified int64) error {

	list_cb := func(ctx context.Context, job *Job) error {

		if job.Status != status {
			return nil
		}

		if job.LastModified > lastmodified {
			return nil
		}

		err := db.RemoveJob(ctx, job)

		if err != nil {
			return err
		}

		return nil
	}

	return db.ListJobs(ctx, list_cb)
}

// ListJobs() iterates through all the jobs in 'db' passing each to 'list_cb'.
func (db *SyncMapDatabase) ListJobs(ctx context.Context, list_cb ListJobsCallback) error {

	var list_err error

	db.sync_map.Range(func(k interface{}, v interface{}) bool {

		job := v.(*Job)
		err := list_cb(ctx, job)

		if err != nil {
			list_err = err
			return false
		}

		return true
	})

	return list_err
}

func (db *SyncMapDatabase) Close(ctx context.Context) error {
	return nil
}
