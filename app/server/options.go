package server

import (
	"context"
	"flag"
	"fmt"
	"net/url"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-offline"
)

type RunOptions struct {
	OfflineDatabaseURI   string
	OfflineQueueMux      map[string]offline.Queue
	AuthenticatorURI     string
	EnableCORS           bool
	CORSAllowedOrigins   []string
	CORSAllowCredentials bool
	Verbose              bool
}

func DeriveRunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "OFFLINE", false)

	if err != nil {
		return nil, fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	ctx := context.Background()

	// START OF put me in a function with well-defined types etc.

	q_mux := make(map[string]offline.Queue)

	for _, kv := range offline_queue_uris {

		job_type := kv.Key()
		offline_uri := kv.Value().(string)

		_, exists := q_mux[job_type]

		if exists {
			return nil, fmt.Errorf("Multiple values for '%s' job type", job_type)
		}

		offline_uri, err := url.QueryUnescape(offline_uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to unescape URI '%s' for job '%s', %w", offline_uri, job_type, err)
		}

		offline_q, err := offline.NewQueue(ctx, offline_uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to instantiate offline queue for '%s', %w", job_type, err)
		}

		q_mux[job_type] = offline_q
	}

	// END OF put me in a function with well-defined types etc.

	opts := &RunOptions{
		OfflineDatabaseURI:   offline_database_uri,
		OfflineQueueMux:      q_mux,
		AuthenticatorURI:     authenticator_uri,
		EnableCORS:           enable_cors,
		CORSAllowedOrigins:   cors_origins,
		CORSAllowCredentials: cors_allow_credentials,
		Verbose:              verbose,
	}

	return opts, nil
}
