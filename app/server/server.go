package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-server/handler"
)

func Run(ctx context.Context, logger *slog.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	opts, err := DeriveRunOptionsFromFlagSet(fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts, logger)
}

func RunWithOptions(ctx context.Context, opts *RunOptions, logger *slog.Logger) error {

	run_opts = opts
	
	handlers := map[string]handler.RouteHandlerFunc{
		"/status":   statusHandlerFunc,
		"/schedule": scheduleHandlerFunc,
	}

	log_logger := slog.NewLogLogger(logger.Handler(), slog.LevelInfo)

	route_handler_opts := &handler.RouteHandlerOptions{
		Handlers: handlers,
		Logger:   log_logger,
	}

	route_handler, err := handler.RouteHandlerWithOptions(route_handler_opts)

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
