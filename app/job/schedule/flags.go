package schedule

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var offline_database_uri string
var offline_queue_uris multi.KeyValueString
var job_id int64

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&offline_database_uri, "offline-database-uri", "", "...")
	fs.Var(&offline_queue_uris, "offline-queue-uri", "...")
	fs.Int64Var(&job_id, "job-id", 0, "...")

	return fs
}
