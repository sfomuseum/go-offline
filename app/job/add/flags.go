package add

import (
	"flag"
	"github.com/sfomuseum/go-flags/flagset"
)

var database_uri string
var instructions string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("offline")

	fs.StringVar(&database_uri, "database-uri", "awsdynamodb://offlinejobs?partition_key=Id&local=true", "")
	fs.StringVar(&instructions, "instructions", "", "")

	return fs
}
