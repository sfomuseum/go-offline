package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/sfomuseum/go-offline/http/api"
)

func statusHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		slog.Error("Failed to set up common configuration", "error", setupCommonError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
	}

	status_handler_opts := &api.JobStatusHandlerOptions{
		OfflineDatabase: offline_db,
		Authenticator:   authenticator,
	}

	status_handler := api.JobStatusHandler(status_handler_opts)
	status_handler = authenticator.WrapHandler(status_handler)

	if run_opts.EnableCORS {
		status_handler = cors_wrapper.Handler(status_handler)
	}

	status_handler = authenticator.WrapHandler(status_handler)
	return status_handler, nil
}

func scheduleHandlerFunc(ctx context.Context) (http.Handler, error) {

	setupCommonOnce.Do(setupCommon)

	if setupCommonError != nil {
		slog.Error("Failed to set up common configuration", "error", setupCommonError)
		return nil, fmt.Errorf("Failed to set up common configuration, %w", setupCommonError)
	}

	schedule_handler_opts := &api.ScheduleJobHandlerOptions{
		OfflineDatabase: offline_db,
		OfflineQueueMux: run_opts.OfflineQueueMux,
		Authenticator:   authenticator,
	}

	schedule_handler := api.ScheduleJobHandler(schedule_handler_opts)
	schedule_handler = authenticator.WrapHandler(schedule_handler)

	if run_opts.EnableCORS {
		schedule_handler = cors_wrapper.Handler(schedule_handler)
	}

	schedule_handler = authenticator.WrapHandler(schedule_handler)
	return schedule_handler, nil
}
