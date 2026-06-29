package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

var (
	cyan   = "\033[36m"
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	reset  = "\033[0m"
)

func statusColor(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return yellow
	case code >= 400 && code < 500:
		return red
	default:
		return cyan
	}
}

var logger = log.New(os.Stdout, "", 0)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)

		method := fmt.Sprintf("%s%s%s", cyan, r.Method, reset)
		path := r.URL.Path
		status := fmt.Sprintf("%s%d%s", statusColor(sw.status), sw.status, reset)
		duration := fmt.Sprintf("%s%v%s", yellow, time.Since(start), reset)
		line := fmt.Sprintf("%s %s %s %s\n", method, path, status, duration)
		_, err := os.Stdout.WriteString(line)
		if err != nil {
			return
		}
	})
}
