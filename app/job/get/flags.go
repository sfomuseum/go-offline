package get

import (
	"flag"
	"github.com/sfomuseum/go-flags/flagset"
)

var database_uri string
var job_id int64

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&database_uri, "database-uri", "", "")
	fs.Int64Var(&job_id, "job-id", 0, "")

	return fs
}
