package server

import (
	"context"
	"fmt"

	"github.com/aaronland/go-http/v4/auth"
	"github.com/rs/cors"
	"github.com/sfomuseum/go-offline"
)

func setupCommon() {

	ctx := context.Background()
	var err error

	offline_db, err = offline.NewDatabase(ctx, run_opts.OfflineDatabaseURI)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to create offline database, %w", err)
		return
	}

	authenticator, err = auth.NewAuthenticator(ctx, run_opts.AuthenticatorURI)

	if err != nil {
		setupCommonError = fmt.Errorf("Failed to create authenticator, %w", err)
		return
	}

	if run_opts.EnableCORS {

		if len(run_opts.CORSAllowedOrigins) == 0 {
			setupCommonError = fmt.Errorf("Missing allowed CORS origin hosts")
			return
		}

		cors_wrapper = cors.New(cors.Options{
			AllowedOrigins:   run_opts.CORSAllowedOrigins,
			AllowCredentials: run_opts.CORSAllowCredentials,
		})
	}

}
