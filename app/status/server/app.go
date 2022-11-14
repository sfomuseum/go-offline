package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-server"
	"github.com/rs/cors"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-offline"
	"github.com/sfomuseum/go-offline/http/api"
	"log"
	"net/http"
)

type RunOptions struct {
	FlagSet       *flag.FlagSet
	Logger        *log.Logger
	EnvFlagPrefix string
}

func Run(ctx context.Context, logger *log.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	opts := &RunOptions{
		FlagSet:       fs,
		Logger:        logger,
		EnvFlagPrefix: "OFFLINE",
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	fs := opts.FlagSet
	logger := opts.Logger

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, opts.EnvFlagPrefix, true)

	if err != nil {
		return fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	offline_db, err := offline.NewDatabase(ctx, offline_database_uri)

	if err != nil {
		return fmt.Errorf("Failed to create offline database, %w", err)
	}

	authenticator, err := auth.NewAuthenticator(ctx, authenticator_uri)

	if err != nil {
		return fmt.Errorf("Failed to create authenticator, %w", err)
	}

	var cors_wrapper *cors.Cors

	if enable_cors {

		if len(cors_origins) == 0 {
			return fmt.Errorf("Missing allowed CORS origin hosts")
		}

		cors_wrapper = cors.New(cors.Options{
			AllowedOrigins:   cors_origins,
			AllowCredentials: cors_allow_credentials,
		})

	}

	status_handler_opts := &api.JobStatusHandlerOptions{
		Database:      offline_db,
		Authenticator: authenticator,
	}

	status_handler := api.JobStatusHandler(status_handler_opts)
	status_handler = authenticator.WrapHandler(status_handler)

	if enable_cors {
		status_handler = cors_wrapper.Handler(status_handler)
	}

	mux := http.NewServeMux()

	mux.Handle(path_status, status_handler)

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		return fmt.Errorf("Failed to create server, %w", err)
	}

	logger.Printf("Listening for requests at %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to serve requests, %w", err)
	}

	return nil
}
