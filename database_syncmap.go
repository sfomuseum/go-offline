package offline

import (
	"context"
	"fmt"
	"sync"
)

type SyncMapDatabase struct {
	Database
	sync_map *sync.Map
}

func init() {
	ctx := context.Background()
	RegisterDatabase(ctx, "syncmap", NewSyncMapDatabase)
}

func NewSyncMapDatabase(ctx context.Context, uri string) (Database, error) {

	sync_map := new(sync.Map)

	db := &SyncMapDatabase{
		sync_map: sync_map,
	}

	return db, nil
}

func (db *SyncMapDatabase) AddJob(ctx context.Context, job *Job) error {
	db.sync_map.Store(job.Id, job)
	return nil
}

func (db *SyncMapDatabase) GetJob(ctx context.Context, job_id int64) (*Job, error) {
	v, ok := db.sync_map.Load(job_id)

	if !ok {
		return nil, fmt.Errorf("Not found")
	}

	return v.(*Job), nil
}

func (db *SyncMapDatabase) UpdateJob(ctx context.Context, job *Job) error {
	db.sync_map.Store(job.Id, job)
	return nil
}

func (db *SyncMapDatabase) RemoveJob(ctx context.Context, job *Job) error {
	db.sync_map.Delete(job.Id)
	return nil
}

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
