package http

import (
	"log/slog"
	"net/http"
)

// LoggerWithRequest assigns informative headers and request details to 'logger'
func LoggerWithRequest(req *http.Request, logger *slog.Logger) *slog.Logger {

	if logger == nil {
		logger = slog.Default()
	}

	logger = logger.With("method", req.Method)
	logger = logger.With("accept", req.Header.Get("Accept"))
	logger = logger.With("path", req.URL.Path)
	logger = logger.With("remote addr", req.RemoteAddr)
	logger = logger.With("user ip", ReadUserIP(req))

	return logger
}

// ReadUserIP returns the value of the 'X-Real-Ip' or 'X-Forwarded-For' headers if present and,
// if not, returns the default "RemoteAddr" value.
func ReadUserIP(req *http.Request) string {

	addr := req.Header.Get("X-Real-Ip")

	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
	}

	if addr == "" {
		addr = req.RemoteAddr
	}

	return addr
}
