package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/prashantv/go_profiling_talk/stats"

	"github.com/varstr/uaparser"
)

const serviceName = "demo-app"

var _hostName = getHost()

// WithStats wraps handlers with stats reporting. It tracks metrics such
// as the number of requests per endpoint, the latency, etc.
func WithStats(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		tags := getStatsTags(r)
		stats.IncCounter("handler.received", tags, 1)

		h(w, r)

		duration := time.Since(start)
		stats.RecordTimer("handler.latency", tags, duration)
	}
}

func getHost() string {
	host, err := os.Hostname()
	if err != nil {
		return ""
	}

	if idx := strings.IndexByte(host, '.'); idx > 0 {
		host = host[:idx]
	}
	return host
}

func getStatsTags(r *http.Request) map[string]string {
	userBrowser, userOS := parseUserAgent(r.UserAgent())
	stats := map[string]string{
		"browser":  userBrowser,
		"os":       userOS,
		"endpoint": filepath.Base(r.URL.Path),
	}
	if _hostName != "" {
		stats["host"] = _hostName
	}
	return stats
}

func parseUserAgent(uaString string) (browser, os string) {
	ua := uaparser.Parse(uaString)

	if ua.Browser != nil {
		browser = ua.Browser.Name
	}
	if ua.OS != nil {
		os = ua.OS.Name
	}

	return browser, os
}
