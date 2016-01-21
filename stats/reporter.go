package stats

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var printStats = flag.Bool("printStats", false, "Print stats to console")

// IncCounter increments a counter.
func IncCounter(name string, tags map[string]string, value int64) {
	name = addTagsToName(name, tags)
	if *printStats {
		fmt.Printf("IncCounter: %v = %v\n", name, value)
	}
}

// UpdateGauge updates a gauge.
func UpdateGauge(name string, tags map[string]string, value int64) {
	name = addTagsToName(name, tags)
	if *printStats {
		fmt.Printf("UpdateGauge: %v = %v\n", name, value)
	}
}

// RecordTimer records a timer.
func RecordTimer(name string, tags map[string]string, d time.Duration) {
	name = addTagsToName(name, tags)
	if *printStats {
		fmt.Printf("RecordTimer: %v = %v\n", name, d)
	}
}

func addTagsToName(name string, tags map[string]string) string {
	// The format we want is: host.endpoint.os.browser
	// if there's no host tag, then we don't use it.
	keyOrder := make([]string, 0, 4)
	if _, ok := tags["host"]; ok {
		keyOrder = append(keyOrder, "host")
	}
	keyOrder = append(keyOrder, "endpoint", "os", "browser")

	parts := make([]string, 1, 5)
	parts[0] = name
	for _, k := range keyOrder {
		v, ok := tags[k]
		if !ok || v == "" {
			parts = append(parts, "no-"+k)
			continue
		}
		parts = append(parts, clean(v))
	}

	return strings.Join(parts, ".")
}

// clean takes a string that may contain special characters, and replaces these
// characters with a '-'.
func clean(value string) string {
	newStr := make([]byte, len(value))
	for i := 0; i < len(value); i++ {
		switch c := value[i]; c {
		case '{', '}', '/', '\\', ':', ' ', '\t', '.':
			newStr[i] = '-'
		default:
			newStr[i] = c
		}
	}
	return string(newStr)
}
