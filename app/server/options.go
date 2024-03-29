package server

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	OfflineDatabaseURI   string
	OfflineQueueURI      string
	AuthenticatorURI     string
	EnableCORS           bool
	CORSAllowedOrigins   []string
	CORSAllowCredentials bool
}

func DeriveRunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "OFFLINE", false)

	if err != nil {
		return nil, fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	opts := &RunOptions{
		OfflineDatabaseURI:   offline_database_uri,
		OfflineQueueURI:      offline_queue_uri,
		AuthenticatorURI:     authenticator_uri,
		EnableCORS:           enable_cors,
		CORSAllowedOrigins:   cors_origins,
		CORSAllowCredentials: cors_allow_credentials,
	}

	return opts, nil
}
