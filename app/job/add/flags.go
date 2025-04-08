package add

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var database_uri string
var creator string
var job_type string
var instructions string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&database_uri, "database-uri", "", "")
	fs.StringVar(&creator, "creator", "", "")
	fs.StringVar(&job_type, "type", "", "")
	fs.StringVar(&instructions, "instructions", "", "")

	return fs
}
