package remove

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var offline_database_uri string
var job_id int64
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&offline_database_uri, "offline-database-uri", "", "A registered sfomuseum/go-offline.Database URI.")
	fs.Int64Var(&job_id, "job-id", 0, "")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
