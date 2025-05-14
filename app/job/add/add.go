package add

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-offline"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	db, err := offline.NewDatabase(ctx, offline_database_uri)

	if err != nil {
		return fmt.Errorf("Failed to create offline database for '%s', %w", offline_database_uri, err)
	}

	job, err := offline.NewJob(ctx, creator, job_type, instructions)

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
