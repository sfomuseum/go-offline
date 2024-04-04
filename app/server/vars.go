package server

import (
	"sync"

	"github.com/rs/cors"
	"github.com/sfomuseum/go-http-auth"
	"github.com/sfomuseum/go-offline"
)

var run_opts *RunOptions

var setupCommonOnce sync.Once
var setupCommonError error

var offline_db offline.Database
var authenticator auth.Authenticator

var cors_wrapper *cors.Cors
