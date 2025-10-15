package server

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"

	"github.com/aaronland/gocloud/runtimevar"
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

func DeriveRunOptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "OFFLINE", false)

	if err != nil {
		return nil, fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	// START OF put me in a function with well-defined types etc.

	if offline_queue_config_uri != "" {

		cfg_str, err := runtimevar.StringVar(ctx, offline_queue_config_uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to derive string value for offline-queue-config-uri, %w", err)
		}

		var offline_cfg map[string]string

		err = json.Unmarshal([]byte(cfg_str), &offline_cfg)

		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal offline-queue-config, %w", err)
		}

		for task, uri := range offline_cfg {

			str_flag := fmt.Sprintf("%s=%s", task, uri)
			err = offline_queue_uris.Set(str_flag)

			if err != nil {
				return nil, fmt.Errorf("Failed to assign offline task flag (%s), %w", task, err)
			}
		}
	}

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
