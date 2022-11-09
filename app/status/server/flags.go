package server

import (
	"flag"
	"github.com/sfomuseum/go-flags/flagset"
)

var offline_database_uri string

var authenticator_uri string

var server_uri string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&offline_database_uri, "offline-database-uri", "", "")

	fs.StringVar(&authenticator_uri, "authenticator-uri", "null://", "")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "")

	return fs
}
