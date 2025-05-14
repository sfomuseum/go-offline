package schedule

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var offline_database_uri string
var offline_queue_uris multi.KeyValueString
var job_id int64
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&offline_database_uri, "offline-database-uri", "", "A registered sfomuseum/go-offline.Database URI.")
	fs.Var(&offline_queue_uris, "offline-queue-uri", "One or more {TASK}={QUEUE_URI} key-value pairs mapping a job type to a registered sfomuseum/go-offline.Queue URI.")
	fs.Int64Var(&job_id, "job-id", 0, "The job ID to schedule.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
