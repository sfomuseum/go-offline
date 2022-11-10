package server

import (
	"flag"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var offline_database_uri string

var authenticator_uri string

var server_uri string

var enable_cors bool

var cors_origins multi.MultiCSVString

var cors_allow_credentials bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&offline_database_uri, "offline-database-uri", "", "")

	fs.StringVar(&authenticator_uri, "authenticator-uri", "null://", "")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "")

	fs.BoolVar(&enable_cors, "enable-cors", true, "A boolean flag to enable CORS headers")

	fs.Var(&cors_origins, "cors-origin", "One or more hosts to restrict CORS support to on the API endpoint.")

	fs.BoolVar(&cors_allow_credentials, "cors-allow-credentials", false, "A boolean flag indicating whether or not to allow credentials headers for CORS requests.")

	return fs
}
