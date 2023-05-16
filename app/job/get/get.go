package get

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-offline"
)

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	db, err := offline.NewDatabase(ctx, database_uri)

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
