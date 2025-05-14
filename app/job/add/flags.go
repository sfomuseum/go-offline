package add

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var offline_database_uri string
var creator string
var job_type string
var instructions string
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&offline_database_uri, "offline-database-uri", "", "A registered sfomuseum/go-offline.Database URI.")
	fs.StringVar(&creator, "creator", "", "")
	fs.StringVar(&job_type, "type", "", "")
	fs.StringVar(&instructions, "instructions", "", "")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
