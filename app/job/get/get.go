package get

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"

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
		return fmt.Errorf("Failed to create offline database, %w", err)
	}

	job, err := db.GetJob(ctx, job_id)

	if err != nil {
		return fmt.Errorf("Failed to get job, %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(&job)

	if err != nil {
		return fmt.Errorf("Failed to decode job, %w", err)
	}

	return nil
}
