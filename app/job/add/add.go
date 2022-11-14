package add

import (
	"context"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-offline"
	"log"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	db, err := offline.NewDatabase(ctx, database_uri)

	if err != nil {
		return fmt.Errorf("Failed to create offline database for '%s', %w", database_uri, err)
	}

	job, err := offline.NewJob(ctx, instructions)

	if err != nil {
		return fmt.Errorf("Failed to create new job, %w", err)
	}

	err = db.AddJob(ctx, job)

	if err != nil {
		return fmt.Errorf("Failed to add job, %w", err)
	}

	fmt.Println(job.Id)
	return nil
}
