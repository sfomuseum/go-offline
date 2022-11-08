package offline

import (
	"context"
	"fmt"
	"gocloud.dev/docstore"
	"io"
	"time"
)

type DocstoreDatabase struct {
	Database
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err := RegisterDatabase(ctx, scheme, NewDocstoreDatabase)

		if err != nil {
			panic(err)
		}
	}
}

func NewDocstoreDatabase(ctx context.Context, uri string) (Database, error) {

	col, err := docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create collection, %w", err)
	}

	db := &DocstoreDatabase{
		collection: col,
	}

	return db, nil
}

func (db *DocstoreDatabase) AddJob(ctx context.Context, job *Job) error {

	err := db.collection.Put(ctx, job)

	if err != nil {
		return fmt.Errorf("Failed to store job, %w", err)
	}

	return nil
}

func (db *DocstoreDatabase) GetJob(ctx context.Context, id int64) (*Job, error) {

	q := db.collection.Query()
	q = q.Where("Id", "=", id)

	iter := q.Get(ctx)
	defer iter.Stop()
	
	var job Job
	err := iter.Next(ctx, &job)

	if err != nil {

		if err == io.EOF {
			return nil, fmt.Errorf("Not found")
		}

		return nil, fmt.Errorf("Failed to retrieve next item in iterator, %w", err)
	}

	return &job, nil
}

func (db *DocstoreDatabase) UpdateJob(ctx context.Context, job *Job) error {

	now := time.Now()
	ts := now.Unix()

	job.LastModified = ts

	return db.collection.Replace(ctx, job)
}

func (db *DocstoreDatabase) RemoveJob(ctx context.Context, job *Job) error {

	return db.collection.Delete(ctx, job)
}
