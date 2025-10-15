package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-http/v3/handlers"
	"github.com/aaronland/go-http/v3/server"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := DeriveRunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	run_opts = opts

	route_handlers := map[string]handlers.RouteHandlerFunc{
		"GET /status":    statusHandlerFunc,
		"POST /schedule": scheduleHandlerFunc,
	}

	route_handler_opts := &handlers.RouteHandlerOptions{
		Handlers: route_handlers,
	}

	route_handler, err := handlers.RouteHandlerWithOptions(route_handler_opts)

	if err != nil {
		return fmt.Errorf("Failed to configure route handler, %w", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", route_handler)

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		return fmt.Errorf("Failed to create server, %w", err)
	}

	slog.Info(fmt.Sprintf("Listening for requests at %s\n", s.Address()))

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to serve requests, %w", err)
	}

	return nil
}
